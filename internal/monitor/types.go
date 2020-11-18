package monitor

import (
	"github.com/pPrecel/raspberrypi-file-monitor/internal/config"
	"github.com/pPrecel/raspberrypi-file-monitor/internal/metric"
	"time"
)

type Files []File

type File struct {
	FileInfo
	Metrics []Metric
}

type FileInfo struct {
	filename string
	delay    time.Duration
}

type Prometheus interface {
	Expose(float64)
	Name() string
}

type Metric struct {
	PrometheusEntry Prometheus
	line, column    int
}

func Parse(entries []config.MonitorEntry) (Files, error) {
	var files Files
	for _, entry := range entries {
		duration, err := time.ParseDuration(entry.TimeDelay)
		if err != nil {
			return nil, err
		}

		file := File{
			FileInfo: FileInfo{
				filename: entry.Filename,
				delay:    duration,
			},
		}
		for _, elem := range entry.Metrics {
			m := Metric{
				PrometheusEntry: metric.New(metric.TypeFromString(elem.Type), elem.Name, elem.Description),
				line:            elem.Line,
				column:          elem.Column,
			}
			file.Metrics = append(file.Metrics, m)
		}
		files = append(files, file)
	}
	return files, nil
}
