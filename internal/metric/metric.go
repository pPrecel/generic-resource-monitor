package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	interact func(float64)
	name     string
}

type metricType int

const (
	Gauge metricType = iota
	Counter
)

func TypeFromString(value string) metricType {
	switch value {
	case "Gauge":
		return Gauge
	case "Counter":
		return Counter
	}
	return TypeFromString("Gauge")
}

func New(metricType metricType, name, help string) *Metric {
	metric := &Metric{
		name: name,
	}
	switch metricType {
	case Gauge:
		prometheusMetric := promauto.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: help,
		})
		metric.interact = prometheusMetric.Set
	case Counter:
		prometheusMetric := promauto.NewCounter(prometheus.CounterOpts{
			Name: name,
			Help: help,
		})
		metric.interact = prometheusMetric.Add
	default:
		metric = New(metricType, name, help)
	}

	return metric
}

func (m *Metric) Expose(value float64) {
	m.interact(value)
}

func (m *Metric) Name() string {
	return m.name
}
