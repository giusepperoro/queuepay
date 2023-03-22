package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func GetConfigFromFile(filename string) (ServiceConfiguration, error) {
	var cfg ServiceConfiguration
	f, err := os.Open(filename)
	if err != nil {
		return cfg, fmt.Errorf("unable to open config file '%s': %w", filename, err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return cfg, fmt.Errorf("unable to read from config file '%s': %w", filename, err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("unable to unmarshall config data: %w", err)
	}
	return cfg, nil
}
