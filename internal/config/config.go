package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func ReadConfig(filename string) (Config, error) {
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
	return config, err
}
