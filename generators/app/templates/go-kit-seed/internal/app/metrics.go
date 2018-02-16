package app

import (
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"
	metrics "github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// NewCounter - creates new metrics counter
func NewCounter(config utils.Configuration, name string, help string) metrics.Counter {
	return prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: config.GetOrgName(),
		Subsystem: config.GetTitle(),
		Name:      name,
		Help:      help,
	}, []string{})
}

// NewHistogram - creates new metrics histogram
func NewHistogram(config utils.Configuration, name string, help string) metrics.Histogram {
	return prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: config.GetOrgName(),
		Subsystem: config.GetTitle(),
		Name:      name,
		Help:      help,
	}, []string{"method", "success"})
}
