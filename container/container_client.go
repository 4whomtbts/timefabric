package container

import (
	"bufio"
	"context"
	"fmt"
	"github.com/4whomtbts/timefabric/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type ContainerClient struct {
	globalCfg *config.TimeFabricConfig
	cfg *config.ContainerConfig
	dockerCtx context.Context
	dockerClient *client.Client
}

func NewContainerClient(globalConfig *config.TimeFabricConfig, containerConfig *config.ContainerConfig) *ContainerClient {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	ctx := context.Background()
	if err != nil {
		log.Error("failed to create docker client")
		panic(err)
	}

	dockerClient.NegotiateAPIVersion(ctx)
	return &ContainerClient{
		globalConfig,
		containerConfig,
		ctx,
		dockerClient,
	}
}

// TODO
func (cc *ContainerClient) buildImage() error {
	return nil
}

func toTimeFabricPortType(ports []types.Port) []Port {
	var newPorts = make([]Port, len(ports))
	for idx, port := range ports {
		newPorts[idx] = Port{
			port.IP,
			port.PrivatePort,
			port.PublicPort,
			port.Type,
		}
	}
	return newPorts
}

func toTimeFabricMountType(mounts []types.MountPoint) []MountPoint {
	var newMounts = make([]MountPoint, len(mounts))
	for idx, mount := range mounts {
		newMounts[idx] = MountPoint{
			mount.Type,
			mount.Name,
			mount.Source,
			mount.Destination,
			mount.Driver,
			mount.Mode,
			mount.RW,
		}
	}
	return newMounts
}

func (cc *ContainerClient) getAllContainers() (string, error) {

	conts, err := cc.dockerClient.ContainerList(cc.dockerCtx, types.ContainerListOptions{})
	if err != nil {
		log.Errorf("failed to fetch container list : [ %s ]", err.Error())
		return "", nil
	}

	var rsConts = make([]Container, len(conts))
	for idx, cont := range conts {
		rsConts[idx] = Container{
			cont.ID,
			cont.Names,
			cont.Image,
			cont.ImageID,
			cont.Command,
			cont.Created,
			toTimeFabricPortType(cont.Ports),
			cont.SizeRw,
			cont.SizeRootFs,
			cont.State,
			cont.Status,
			toTimeFabricMountType(cont.Mounts),
		}
	}

	return "", nil
}

func (cc *ContainerClient) pullImage(fromImage string, tag string) error {

	if len(fromImage) == 0 || len(tag) == 0 {
		log.Error("Empty fromImage or tag is provided")
		return fmt.Errorf("fromImage or tag can't be empty")
	}

	fullImageName := fmt.Sprintf("%s:%s", fromImage, tag)
	log.Infof("Execute command [ docker pull %s ]", fullImageName)

	cmd := exec.Command("docker", "pull", fullImageName)
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()

	if err != nil {
		msg := "Failed to open pipe for command"
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(msg)
		return fmt.Errorf(msg)
	}

	err = cmd.Start()

	if err != nil {
		msg := "Failed to start pulling"
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(msg)
		return fmt.Errorf(msg)
	}

	stdoutScanner := bufio.NewScanner(stdout)
	stdoutScanner.Split(bufio.ScanLines)
	for stdoutScanner.Scan() {
		m := stdoutScanner.Text()
		fmt.Println(m)
	}

	stderrScanner := bufio.NewScanner(stderr)
	stderrScanner.Split(bufio.ScanLines)
	for stderrScanner.Scan() {
		m := stderrScanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()

	if err != nil {
		msg := "Failed to pull Image"
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(msg)
		return fmt.Errorf(msg)
	}

	log.Infof("Successfully pulled image %s !", fullImageName)
	return nil
}
