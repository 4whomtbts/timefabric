package container

import (
	"fmt"
	"strconv"
)

/*
	Maybe.. number granted options
	such as "shm-size, unlimit memlock, ..etc" needs to be handled differently
	Node should check and validate if granted numbers are reasonable and supportable.

var SUPPORTED_CUSTOMIZABLE_CONT_OPTS []string
SUPPORTED_CUSTOMIZABLE_CONT_OPTS {
	"--cap-add", "--ipc", "--shm-size", "--unlimit memlock", "--unlimit stack"
}

*/
/*
	 It extracts the options that timefabric allows to customize.
	 Then, It returns the extracted options as result
 */
func filterNonCustomizableOption(customOption string) string {
	/*
	result := ""
	for idx, val := range  {

	}
	 */
	return ""
}

func buildBindMountMappingOption(sourcePath, targetPath string) string {
	return fmt.Sprintf("--mount type=bind,source=\"%s\",target=%s ", sourcePath, targetPath)
}

func buildVolumeMappingOption(sourcePath, targetPath string) string {
	return fmt.Sprintf("--mount type=volume,source=\"%s\",target=%s ", sourcePath, targetPath)
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
	return gpuOption
}

func buildRunCommand(newContainer ContainerCreate) string {
	return fmt.Sprintf("docker ")
}
