package container

import (
	"fmt"
	"github.com/4whomtbts/timefabric/common"
	"github.com/4whomtbts/timefabric/config"
	"strconv"
	"strings"
)

/*
	Maybe.. number granted options
	such as "shm-size, unlimit memlock, ..etc" needs to be handled differently
	Node should check and validate if granted numbers are reasonable and supportable.

	It would be more easy to test, if this constant slice is ordered ascending
*/
var SUPPORTED_CUSTOMIZABLE_CONT_OPTS = [...]string {
	"--cap-add", "--ipc", "--shm-size", "--unlimit memlock", "--unlimit stack",
}

func consumeUntilEndToken(str *string, base []string, idx int) string {
	for i := idx; i < len(*str); i++ {
		curr := string((*str)[i])
		base = append(base, curr)
		if curr == " " {
			break
		}
		if i == len(*str) - 1 {
			base = append(base, " ")
		}
	}
	return strings.Join(base, "")
}

/*
	 It extracts the options that timefabric allows to customize.
	 Then, It returns the extracted options as result
*/
func filterNonCustomizableOption(customOption string) string {

	result := []string {}

	for _, opt := range SUPPORTED_CUSTOMIZABLE_CONT_OPTS {
		optionStr := []string {}
		idx := strings.Index(customOption, opt)

		if idx == -1 {
			continue
		}

		for i := 0; i < len(opt); i++ {
			optionStr = append(optionStr, string(customOption[idx + i]))
		}
		idx += len(opt)

		for j := idx; j < len(customOption); j++ {
			curr := string(customOption[j])
			optionStr = append(optionStr, curr)
			if curr == "=" || curr == " "{
				result = append(result, consumeUntilEndToken(&customOption, optionStr, j + 1))
				break
			}
		}
	}

	return strings.Join(result, "")
}

func resolveMappingOption(opt, append string) ([]string, []string){
	mappings := strings.Split(opt, ",")
	srcs := make([]string, len(mappings))
	dests := make([]string, len(mappings))

	for i, mapping := range mappings {
		splitted := strings.Split(mapping, ":")

		if len(splitted) == 1 {
			panic(fmt.Sprintf("invalid mapping option %s is provided", mapping))
		}

		src := splitted[0]
		dest := splitted[1]

		if src == "" || dest == "" {
			panic("invalid src or dest is provided")
		}

		if append != "" {
			src = splitted[0]  + fmt.Sprintf("/%s", append)
		}
		srcs[i] = src
		dests[i] = dest
	}
	return srcs, dests
}

func buildBindMountMappingOption(sourcePath, targetPath string) string {
	return fmt.Sprintf("--mount type=bind,source=\"%s\",target=%s ", sourcePath, targetPath)
}

func buildVolumeMappingOption(sourcePath, targetPath string) string {
	return fmt.Sprintf("--mount type=volume,source=\"%s\",target=%s ", sourcePath, targetPath)
}

func buildMountMappingOption(storageType common.StorageType, srcs, dests []string) string {
	result := ""
	for i := 0; i < len(srcs); i++ {
		switch strings.ToLower(string(storageType)) {
		case common.LOCAL:
			result = result + buildVolumeMappingOption(srcs[i], dests[i])
			break
		case common.NAS:
			result = result + buildBindMountMappingOption(srcs[i], dests[i])
			break
		default:
			panic(fmt.Sprintf("invalid storage option %s is provided", storageType))
		}
	}
	return result
}

func buildStaticMappingOption(storageType common.StorageType, staticMappingOpt string) string {
	srcs, dests := resolveMappingOption(staticMappingOpt, "")
	return buildMountMappingOption(storageType, srcs, dests)
}

func buildDynamicMappingOption(storageType common.StorageType, opt string, append string) string {
	srcs, dests := resolveMappingOption(opt, append)
	return buildMountMappingOption(storageType, srcs, dests)
}

func buildMountMappingOptions(config config.StorageConfig, append string) string {

	if append == "" {
		panic("append couldn't be empty")
	}

	staticOps := buildStaticMappingOption(config.StorageType, config.StaticMapping)
	dynamicOps := buildDynamicMappingOption(config.StorageType, config.DynamicMapping, append)
	return staticOps + dynamicOps
}

func buildPortMappingOption(ports []PortCreate) string {
	portOption := ""
	for _, port := range ports {
		curr := "-p " + strconv.Itoa(port.From) + ":" + strconv.Itoa(port.To) + " "
		portOption += curr
	}
	return portOption
}

func buildGpuAllocOption(gpuSlots []int) string {

	if len(gpuSlots) == 0 {
		return ""
	}

	gpuOption := "--gpus device="
	for idx, gpu := range gpuSlots {
		if idx != 0 {
			gpuOption += ","
		}
		gpuOption += strconv.Itoa(gpu)
	}
	gpuOption += " "
	return gpuOption
}

func buildRunCommand(cont ContainerCreate, storageConfig config.StorageConfig, append string) string {
	return fmt.Sprintf("docker run -d -it --ipc=host %s%s%s%s--name test %s",
			filterNonCustomizableOption(cont.CustomOption),
			buildGpuAllocOption(cont.GpuSlots),
			buildMountMappingOptions(storageConfig, append),
			buildPortMappingOption(cont.Ports),
			cont.ImageName)
}
