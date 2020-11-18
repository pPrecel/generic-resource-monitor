package monitor

import (
	"github.com/pPrecel/raspberrypi-file-monitor/internal/scanner"
	"github.com/pkg/errors"
	"time"
)

type Monitor struct {
	Metrics      Files
	readFile     func(string) ([]string, error)
	extractFloat func([]string, int, int) (float64, error)
}

func New(files Files) *Monitor {
	return &Monitor{
		Metrics:      files,
		readFile:     scanner.ReadLines,
		extractFloat: scanner.ExtractFloat,
	}
}

type Channel struct {
	Error error
	Entry *Metric
	Float float64
}

func (m *Monitor) FireEntries() chan Channel {
	channel := make(chan Channel)
	for _, file := range m.Metrics {
		go func(channel chan Channel, file File) {
			for {
				m.entryFunc(channel, file)
				time.Sleep(file.delay)
			}
		}(channel, file)
	}
	return channel
}

const (
	readingFileErrorFormat     = "while reading file %s"
	extractingValueErrorFormat = "while extracting value from file %s"
)

func (m *Monitor) entryFunc(channel chan Channel, entry File) {
	lines, err := m.readFile(entry.filename)
	if err != nil {
		channel <- Channel{
			Error: errors.Wrapf(err, readingFileErrorFormat, entry.filename),
		}
		return
	}

	for _, metric := range entry.Metrics {
		float, err := m.extractFloat(lines, metric.line, metric.column)
		if err != nil {
			channel <- Channel{
				Error: errors.Wrapf(err, extractingValueErrorFormat, entry.filename),
			}
			return
		}
		channel <- Channel{
			Entry: &metric,
			Float: float,
		}
	}
}
