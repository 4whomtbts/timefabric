package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildBindMountMappingOption(t *testing.T) {
	sourcePath := "/294t/nfs/foodir"
	targetPath := "/home/irteam/dir"
	result := buildBindMountMappingOption(sourcePath, targetPath)
	assert.Equal(t, "--mount type=bind,source=\"/294t/nfs/foodir\",target=/home/irteam/dir ", result)
}

func Test_buildBindVolumeMappingOption(t *testing.T) {
	sourcePath := "/294t/nfs/foodir"
	targetPath := "/home/irteam/dir"
	result := buildVolumeMappingOption(sourcePath, targetPath)
	assert.Equal(t, "--mount type=volume,source=\"/294t/nfs/foodir\",target=/home/irteam/dir ", result)
}

func Test_buildPortMappingOption(t *testing.T) {
	portCreates := []PortCreate{
		{"ssh", 22, 22,},
		{"jupyter", 8081, 8081,},
		{"prometheus", 9100, 9100},
	}
	builtPortOption := buildPortMappingOption(portCreates)
	assert.Equal(t, "-p 22:22 -p 8081:8081 -p 9100:9100 ", builtPortOption)
}

func Test_buildGpuAllocOption(t *testing.T) {
	gpuSlots := []int{0, 1, 4}
	buildGpuOption := buildGpuAllocOption(gpuSlots)
	assert.Equal(t, "--gpus device=0,1,4", buildGpuOption)
}

func Test_buildGpuAllocOption_WHEN_emptySliceProvided_THEN_returnEmptyString(t *testing.T) {
	gpuSlots := []int{}
	shouldBeemptyString := buildGpuAllocOption(gpuSlots)
	assert.Equal(t, "", shouldBeemptyString)
}