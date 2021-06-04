package config


type TimeFabricConfig struct {
	MasterServer string `yaml:"masterServer"`
	MasterAllocation bool `yaml:"masterAlloc"`
	NetworkConfig NetworkConfig `yaml:"network"`
	StorageConfig StorageConfig `yaml:"storage"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
	Secret string `yaml:"secret"`
}

type ContainerConfig struct {
	ApiHost string `yaml:"apiHost"`
}

type NetworkConfig struct {

	/*
	 * NetworkConfig could be either nat or internal
	 * If two or more timefabric clients located behind same firewall or router
	 * They must have been allocated mutually different port range
	 * So, timefabric should identify range of port that current client is allocated to avoid
	 * allocating unreachable port for current client(Because NAT configuration doesn't allow traffic to flow in).
	 */
	NetworkType string `yaml:"type"`
	PortRange string `yaml:"portRange"` // port ranges to be allocated for new container
	ExcludedPorts string `yaml:"excludedPorts"` // ports to be excluded when mapping ports for new container
	DatabaseConfig DatabaseConfig `yaml:"database"`
}

type StorageConfig struct {

	/*
	 * StorageType could be either local or nas
	 * when 'local': persistent storage located on specified location of container host machine
	 * when 'nas': persistent storage is nas. bind mapping container with specified nas directory
	 */
	StorageType string `yaml:"type"`
	/*
	 * directory mapping for docker volume.
	 * It should be form of 'HOST_DIR:CONTAINER_DIR'
	 * examples : /nfs/mydir:/home/irteam/mystorage,/nfs/yourdir:/home/irteam/yourstorage
	 */
	Directories string `yaml:"dir"`

}

type DatabaseConfig struct {
	/*
	 * Database url form of http(s)://HOST:PORT?PARAMS
	 */
	Url string `yaml:"url"`
	/*
	 * User of database connection
	 */
	User string `yaml:"user"`
	Secret string  `yaml:"secret"`
}

