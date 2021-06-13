package container

import (
	"fmt"
	"github.com/4whomtbts/timefabric/common"
	"github.com/4whomtbts/timefabric/config"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

/*
NewContainerClient(&config.TimeFabricConfig{
		StorageConfig: config.StorageConfig{
			StorageType: "nas",
			DynamicMapping: "",
		},
	},
	&config.ContainerConfig{
		ApiHost: "http://localhost:2375",
	})
 */

func TestResolveMappingOption(t *testing.T) {
	srcs, dirs :=
		resolveMappingOption(
			"/a/b/mypath:/home/irteam,/a/b/yourpath:/home/aiuser", "mystorage")
	expectedSrcs := []string {"/a/b/mypath/mystorage", "/a/b/yourpath/mystorage"}
	assert.True(t, reflect.DeepEqual(expectedSrcs, srcs))
	expectedDirs := []string {"/home/irteam", "/home/aiuser"}
	assert.True(t, reflect.DeepEqual(expectedDirs, dirs))
}

func TestResolveMappingOption_WHEN_invalid_mapping_option_is_provided_THEN_PANIC(t *testing.T) {
	assert.Panics(t, func() { resolveMappingOption(
		"something-invalid,can't-reach-here", "mystorage") })
}

func TestResolveMappingOption_WHEN_invalid_mapping_option_is_provided_THEN_PANIC2(t *testing.T) {
	assert.Panics(t, func() { resolveMappingOption(
		":invalid-dest,invalid-src:", "mystorage") })
}

func TestBuildMountMappingOption_WHEN_local_Option_is_provided_THEN_returns_volumeMountOption(t *testing.T) {
	srcs, dirs :=
		resolveMappingOption(
			"/a/b/mypath:/home/irteam,/a/b/yourpath:/home/aiuser", "mystorage")

	result := buildMountMappingOption(common.LOCAL, srcs, dirs)
	assert.Equal(t,
		"--mount type=volume,source=\"/a/b/mypath/mystorage\",target=/home/irteam " +
			"--mount type=volume,source=\"/a/b/yourpath/mystorage\",target=/home/aiuser ", result)
}

func TestBuildMountMappingOption_WHEN_nas_Option_is_provided_THEN_returns_bindMountOption(t *testing.T) {
	srcs, dirs :=
		resolveMappingOption(
			"/a/b/mypath:/home/irteam,/a/b/yourpath:/home/aiuser", "mystorage")

	result := buildMountMappingOption(common.NAS, srcs, dirs)
	assert.Equal(t,
		"--mount type=bind,source=\"/a/b/mypath/mystorage\",target=/home/irteam " +
			"--mount type=bind,source=\"/a/b/yourpath/mystorage\",target=/home/aiuser ", result)
}

func TestBuildMountMappingOption_WHEN_unsupported_type_is_provided_THEN_returns_PANIC(t *testing.T) {
	newCli := NewContainerClient(&config.TimeFabricConfig{
		StorageConfig: config.StorageConfig{
			StorageType: "illegal!!!!",
		},
	},
		&config.ContainerConfig{
			ApiHost: "http://localhost:2375",
		})

	goodSrcs, goodDirs :=
		resolveMappingOption(
			"/a/b/mypath:/home/irteam,/a/b/yourpath:/home/aiuser", "mystorage")

	assert.Panics(t, func() {buildMountMappingOption(newCli.globalCfg.StorageConfig.StorageType, goodSrcs, goodDirs)})
}

func TestBuildBindMountMappingOption(t *testing.T) {
	sourcePath := "/a/b/foodir"
	targetPath := "/home/irteam/dir"
	result := buildBindMountMappingOption(sourcePath, targetPath)
	assert.Equal(t, "--mount type=bind,source=\"/a/b/foodir\",target=/home/irteam/dir ", result)
}

func TestBuildBindVolumeMappingOption(t *testing.T) {
	sourcePath := "/a/b/foodir"
	targetPath := "/home/irteam/dir"
	result := buildVolumeMappingOption(sourcePath, targetPath)
	assert.Equal(t, "--mount type=volume,source=\"/a/b/foodir\",target=/home/irteam/dir ", result)
}

func TestStaticMappingOption_THEN_Returns_BuiltMountOption(t *testing.T) {
	result := buildStaticMappingOption(common.LOCAL, "/a/b/mypath:/home/irteam")
	assert.Equal(t, "--mount type=volume,source=\"/a/b/mypath\",target=/home/irteam ", result)
}

func TestDynamicMappingOption_THEN_Returns_BuiltMountOption(t *testing.T) {
	result := buildStaticMappingOption(common.NAS, "/a/b/mypath:/home/irteam")
	assert.Equal(t, "--mount type=bind,source=\"/a/b/mypath\",target=/home/irteam ", result)
}

func TestBuildMountMappingOptions_WHEN_VALID_THEN_Returns_built_options(t *testing.T) {
	conifg := config.StorageConfig{
		StorageType: common.LOCAL,
		StaticMapping: "/a/b/mypath:/home/irteam,/a/b/yourpath:/home/aiuser",
		DynamicMapping: "/a/b:/x/y,/c/d:/t/s",
	}
	result := buildMountMappingOptions(conifg, "amazing")

	assert.Equal(t,
		"--mount type=volume,source=\"/a/b/mypath\",target=/home/irteam " +
		"--mount type=volume,source=\"/a/b/yourpath\",target=/home/aiuser " +
		"--mount type=volume,source=\"/a/b/amazing\",target=/x/y " +
		"--mount type=volume,source=\"/c/d/amazing\",target=/t/s ", result)
}

func TestBuildMountMappingOptions_WHEN_empty_append_is_provided_THEN_Panic(t *testing.T) {
	conifg := config.StorageConfig{
		StorageType: common.LOCAL,
		StaticMapping: "/a/b/mypath:/home/irteam,/a/b/yourpath:/home/aiuser",
		DynamicMapping: "/a/b:/x/y,/c/d:/t/s",
	}
	assert.Panics(t, func() { buildMountMappingOptions(conifg, "") })
}

func TestBuildPortMappingOption(t *testing.T) {
	portCreates := []PortCreate{
		{"ssh", 22, 22,},
		{"jupyter", 8081, 8081,},
		{"prometheus", 9100, 9100},
	}
	builtPortOption := buildPortMappingOption(portCreates)
	assert.Equal(t, "-p 22:22 -p 8081:8081 -p 9100:9100 ", builtPortOption)
}

func TestBuildGpuAllocOption(t *testing.T) {
	gpuSlots := []int{0, 1, 4}
	buildGpuOption := buildGpuAllocOption(gpuSlots)
	assert.Equal(t, "--gpus device=0,1,4 ", buildGpuOption)
}

func TestBuildGpuAllocOption_WHEN_emptySliceProvided_THEN_returnEmptyString(t *testing.T) {
	gpuSlots := []int{}
	shouldBeemptyString := buildGpuAllocOption(gpuSlots)
	assert.Equal(t, "", shouldBeemptyString)
}

func TestFilterNonCustomizableOption_Extracts_and_returns_only_supported_options(t *testing.T) {
	SHM_OPT := "--shm-size=5000mb "
	CAP_OPT := "--cap-add=SYS_ADMIN "
	MEM_LOCK_OPT := "--unlimit memlock 500mb "
	result :=
		filterNonCustomizableOption(
			fmt.Sprintf("--invalid-opt 100 --not-good=bad %s%s%s",
				SHM_OPT, CAP_OPT, MEM_LOCK_OPT))
	assert.True(t, strings.Contains(result, SHM_OPT))
	assert.True(t, strings.Contains(result, CAP_OPT))
	assert.True(t, strings.Contains(result, MEM_LOCK_OPT))
	assert.Equal(t, len(SHM_OPT) + len(CAP_OPT) + len(MEM_LOCK_OPT), len(result))
}

func TestBuildRunCommand(t *testing.T) {
	userName := "aiuser"
	imageName := "ubuntu:18.04"
	cont := ContainerCreate{
		imageName,
		[]int {0, 1, 2},
		[]PortCreate{
			{ "ssh", 22, 22, },
			{ "jupyter", 8081, 8081 },
		},
		" --cap-add=SYS_ADMIN --shm-size=5000mb",
		userName,
	}
	cmd := buildRunCommand(cont, config.StorageConfig{
		StorageType: common.LOCAL,
		StaticMapping: "/a/b:/c/d",
		DynamicMapping: "/s/t:/u/v",
	}, userName)

	assert.Equal(t,
		fmt.Sprintf("docker run -d -it --ipc=host --cap-add=SYS_ADMIN --shm-size=5000mb --gpus device=0,1,2 " +
		"--mount type=volume,source=\"/a/b\",target=/c/d " +
		"--mount type=volume,source=\"/s/t/%s\",target=/u/v " +
		"-p 22:22 -p 8081:8081 --name test %s", userName, imageName), cmd)
}