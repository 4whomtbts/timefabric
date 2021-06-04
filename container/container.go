package container

type Container struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Image string `yaml:"image,omitempty" json:"image,omitempty"`

}

type Port struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	HostPort int `yaml:"hostPort,omitempty" json:"hostPort,omitempty"`
	ContainerPort int `yaml:"containerPort,omitempty" json:"containerPort,omitempty"`
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
}