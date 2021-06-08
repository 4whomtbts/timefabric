package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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