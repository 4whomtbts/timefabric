package container

import (
	"fmt"
	"github.com/4whomtbts/timefabric/common"
	"strconv"
)

/*
	@ userName : The id when user type to sign in timeFabric console site
	@ mappedPath : The path volume is mapped in container (ie. /home/irteam/foo
 */
func buildStorageMappingOption(storageType common.StorageMappingType, userName string, mappedPath string) string {
	return ""
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
