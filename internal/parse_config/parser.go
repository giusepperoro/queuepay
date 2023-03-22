package parse_config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type YamlConfig struct {
	PostgresConnectUrl string
}

func Simple() YamlConfig {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	yamlConfig := YamlConfig{}
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		log.Fatal(err)
	}
	return yamlConfig
}
