package main

import (
	"bytes"
	"fmt"
	"github.com/pPrecel/raspberrypi-file-monitor/internal/config"
	"github.com/pPrecel/raspberrypi-file-monitor/internal/monitor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

const defaultFilename = "/app/config.yaml"
const configFilenameEnvKey = "CONFIG_FILENAME"

func main() {
	filename := defaultFilename
	if newFilename := os.Getenv(configFilenameEnvKey); newFilename != "" {
		filename = newFilename
	}

	log.Printf("Read configuration from: %s\n", filename)
	serverConfig, err := config.ReadConfig(filename)
	failOnError(err)

	err = Log(serverConfig)
	failOnError(err)

	log.Info("Fire Up all entries")
	monitor.FireUpEntries(serverConfig.Entries)

	log.Info("Start server")
	http.Handle(serverConfig.ServerAddress, promhttp.Handler())
	http.ListenAndServe(serverConfig.ServerPort, nil)
}

const bannerInfo = "Program Configuration:\n"
const elemFormat = "%s\n"

func Log(config config.Config) error {
	b, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	log.Println(bannerInfo)
	bytes := bytes.Split(b, []byte("\n"))
	for _, line := range bytes {
		log.Println(fmt.Sprintf(elemFormat, string(line)))
	}
	return nil
}

func failOnError(err error) {
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
