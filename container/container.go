package container

import "github.com/docker/docker/api/types/mount"

type Container struct {
	Id string `json:"Id,omitempty"`
	Names []string `json:"Names"`
	Image string `json:"Image"`
	ImageId string `json:"ImageID"`
	Command string `json:"Command"`
	Created int64 `json:"Created"`
	Ports []Port `json:"Ports"`
	SizeRW int64 `json:"SizeRW"`
	SizeRootFS int64 `json:"SizeRootFS"`
	State string `json:"State"`
	Status string `json:"Status"`
	Mounts []MountPoint
}

type ContainerCreate struct {
	ImageName string `json:"image_name"`
	GpuSlots []int `json:"gpu_slots"`
	Ports []PortCreate `json:"ports"`
	CustomOption string `json:"custom_option"`
	UserName string `json:"user_name"`
}

type MountPoint struct {
	Type        mount.Type `json:",omitempty"`
	Name        string     `json:",omitempty"`
	Source      string	`json:"Source"`
	Destination string	`json:"Destination"`
	Driver      string  `json:",omitempty"`
	Mode        string  `json:"Mode"`
	RW          bool	`json:"RW"`
}

type Port struct {
	IP string `yaml:"IP,omitempty" json:"IP,omitempty"`
	PrivatePort uint16 `yaml:"private_port,omitempty" json:"PrivatePort,omitempty"`
	PublicPort uint16 `yaml:"public_port,omitempty" json:"PublicPort,omitempty"`
	Type string `yaml:"type,omitempty" json:"Type,omitempty"`
}

type PortCreate struct {
	MappingName string `json:"mapping_name"`
	From int `json:"from"`
	To int `json:"to"`
}