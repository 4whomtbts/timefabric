package container

import (
	"fmt"
	"github.com/4whomtbts/timefabric/config"
	"github.com/go-resty/resty/v2"
	"os/exec"
)

type ContainerClient struct {
	cfg *config.ContainerConfig
}

func NewContainerClient(containerConfig *config.ContainerConfig) *ContainerClient {
	return &ContainerClient{
			containerConfig,
	}
}

func (cc *ContainerClient) buildImage() error {}

func (cc *ContainerClient) pullImage(fromImage string, tag string) error {}