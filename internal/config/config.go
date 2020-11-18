package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func ReadAndDefault(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return defaultConfig(config), nil
}

const (
	defaultPort      = ":4040"
	defaultAddress   = "/metrics"
	defaultTimeDelay = "5s"
)

func defaultConfig(config Config) Config {
	if config.ServerPort == "" {
		config.ServerPort = defaultPort
	}
	if config.ServerAddress == "" {
		config.ServerAddress = defaultAddress
	}

	for i, entry := range config.Entries {
		if entry.FileInfo.TimeDelay == "" {
			config.Entries[i].TimeDelay = defaultTimeDelay
		}
	}
	return config
}
