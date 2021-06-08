package main

import (
	"fmt"
	"github.com/4whomtbts/timefabric/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func main() {

	configFile, _ := filepath.Abs("example.yaml")
	yamlFile, err := ioutil.ReadFile(configFile)
	var config config.TimeFabricConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Print(config)
}
