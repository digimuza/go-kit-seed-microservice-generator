package app

import (
	"net/http"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/mock"

	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc/reflection"

	"fmt"
	"net"
	"os"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/endpoints"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/models"
	pb "dev.adeoweb.biz/pas/<%= appName %>/pkg/pb/buckets"
	bkservice "dev.adeoweb.biz/pas/<%= appName %>/pkg/service"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/transport"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"

	metrics "github.com/go-kit/kit/metrics"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Start - starts microservice
func Start(config utils.Configuration) {
	var (
		logger  = NewLogger(config)
		db      = NewDBConnection(config, logger)
		tracing = NewTracing(config, logger)
	)

	store, err := models.NewBucketStore(db)
	if err != nil {
		level.Error(logger).Log("transport", "grpc", "err", err.Error())
		os.Exit(1)
	}

	mockData := mock.NewMock()
	pubKeyStore := mock.NewPublicKeyStore(mockData)

	{
		defer db.Close()
		defer tracing.Close()
	}

	var created, quered metrics.Counter
	{
		// Business-level metrics.
		created = NewCounter(
			config,
			"events_created",
			"Total count of events created via the Create method.",
		)
		quered = NewCounter(
			config,
			"events_quered",
			"Total count of events quered via the GetUserEvents method.",
		)
	}
	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = NewHistogram(
			config,
			"request_duration_seconds",
			"Request duration in seconds.",
		)
	}

	var (
		service    = bkservice.NewService(store, pubKeyStore, logger, created, quered)
		controller = endpoints.NewController(service, logger, duration, tracing.OpenTracer())
	)

	var g run.Group
	{
		port := fmt.Sprintf(":%s", config.GetPrometheusMetricsPort())
		metricsListener, err := net.Listen("tcp", port)
		http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
		if err != nil {
			logger.Log("transport", "metrics/HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "metrics/HTTP", "addr", port)
			return http.Serve(metricsListener, http.DefaultServeMux)
		}, func(error) {
			metricsListener.Close()
		})
	}
	{

		grpcServer, err := transport.NewGRPCServer(config)
		if err != nil {
			level.Error(logger).Log("transport", "grpc", "err", err.Error())
			os.Exit(1)
		}
		server := transport.NewServer(controller, tracing.OpenTracer(), logger)
		pb.RegisterBucketsServiceServer(grpcServer, server)
		reflection.Register(grpcServer)

		address := formatServiceURL(config)
		grpcListener, err := net.Listen("tcp", address)
		if err != nil {
			level.Error(logger).Log("transport", "grpc", "during", "Listen", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			level.Info(logger).Log("transport", "grpc", "addr", address)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}

	logger.Log("exit", g.Run())
}

func formatServiceURL(config utils.Configuration) string {
	return fmt.Sprintf(
		"%s:%s",
		config.GetAppName(),
		config.GetPort(),
	)
}
