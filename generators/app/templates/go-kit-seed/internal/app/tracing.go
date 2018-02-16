package app

import (
	"fmt"
	"os"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

// Tracing - object that contains tracing
type Tracing interface {
	OpenTracer() stdopentracing.Tracer
	Close()
}

type tracing struct {
	collector zipkin.Collector
	tracer    stdopentracing.Tracer
}

func (t tracing) OpenTracer() stdopentracing.Tracer {
	return t.tracer
}
func (t tracing) Close() {
	t.collector.Close()
}

// NewTracing - creates tracing
func NewTracing(config utils.Configuration, logger log.Logger) Tracing {
	var (
		zipkinURL   = formatZipkinURL(config)
		debug       = false
		hostPort    = fmt.Sprintf("%s:%s", config.GetAppName(), config.GetPort())
		serviceName = config.GetTitle()
	)

	logger.Log("tracer", "Zipkin", "URL", zipkinURL)

	var collector zipkin.Collector
	{
		var err error
		collector, err = zipkin.NewHTTPCollector(zipkinURL)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
	}

	var tracer stdopentracing.Tracer
	{
		var err error
		recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)
		tracer, err = zipkin.NewTracer(recorder)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
	}
	return tracing{collector, tracer}
}

func formatZipkinURL(config utils.Configuration) string {
	return fmt.Sprintf(
		"http://%s:%s/api/v1/spans",
		config.GetZipkinHost(),
		config.GetZipkinPort(),
	)
}
