package monitor

import (
	"github.com/pPrecel/raspberrypi-file-monitor/internal/config"
	"github.com/pPrecel/raspberrypi-file-monitor/internal/scanner"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	errorFormat = "while reading file %s: %s"
	infoFormat  = "value from file: %s is: %v"
)

var (
	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func FireUpEntries(entries []config.MonitorEntry) error {
	for _, entry := range entries {
		go func(entry config.MonitorEntry) {
			opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
				Name: entry.Name,
				Help: entry.Description,
			})
			for {
				time.Sleep(time.Duration(entry.Timestamp) * time.Second)

				val, err := scanner.ExtractVal(entry.Filename, entry.LineNumber, entry.ElemNumber)
				if err != nil {
					log.Errorf(errorFormat, entry.Filename, err)
				}
				log.Infof(infoFormat, entry.Filename, val)
				opsProcessed.Set(val)
			}
		}(entry)
	}
	return nil
}
