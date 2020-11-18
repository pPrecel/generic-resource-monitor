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
	serverConfig, err := config.ReadAndDefault(filename)
	failOnError(err)

	err = Log(serverConfig)
	failOnError(err)

	log.Info("Fire Up all entries")
	entries, err := monitor.Parse(serverConfig.Entries)
	failOnError(err)

	monitor := monitor.New(entries)
	channel := monitor.FireEntries()

	log.Info("Start server")
	http.Handle(serverConfig.ServerAddress, promhttp.Handler())
	go http.ListenAndServe(serverConfig.ServerPort, nil)

	for {
		value := <-channel
		if value.Error != nil {
			log.Error(value.Error)
		} else if value.Entry != nil {
			value.Entry.PrometheusEntry.Expose(value.Float)
			log.Infof("[%s] read value: %g",
				value.Entry.PrometheusEntry.Name(), value.Float)
		}
	}
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
