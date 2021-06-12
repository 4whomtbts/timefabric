package container

import (
	_ "fmt"
	"github.com/4whomtbts/timefabric/config"
)

var TEST_DOCKER_API_HOST string = "http://localhost:2375"

var testClient *ContainerClient = NewContainerClient(
	&config.TimeFabricConfig{},
	&config.ContainerConfig{
TEST_DOCKER_API_HOST,
})

/*
func Test_pullingImage(t *testing.T) {
	testClient.pullImage("centos", "latest")
//	assert.Nil(t, err)
}

func Test_getAllContainers(t *testing.T) {
	_, err := testClient.getAllContainers()
	assert.Nil(t, err)
}
*/