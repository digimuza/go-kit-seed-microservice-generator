package main

import (
	"fmt"
	<% if (http) { %>
	"net/http"	
	<% } %>
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	<% if (grpc) { %>
	"net"
	"google.golang.org/grpc/reflection"
	"<%= org %>/<%= appName %>/pkg/pb"
	<% } %>
)

func main() {

	<% if (http) { %>
	httpListenPort := fmt.Sprintf(":%d",80)
	<% } %>
	<% if (grpc) { %>
	grpcListenPort := fmt.Sprintf(":%d",443)
	<% } %>	
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var s <%= serviceCamelCase %>Interface
	{
		s = New<%= serviceCamelCase %>()
		s = NewLoggingMiddleware(s, logger)
	}

	<% if (http) { %>
	var h http.Handler
	{
		h = NewHTTPHandler(s, logger)
	}
	<% } %>
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	
	<% if (http) { %>
	go func() {
		logger.Log("transport", "HTTP", "addr", httpListenPort)
		errs <- http.ListenAndServe(httpListenPort , h)
	}()
	<% } %>

	<% if (grpc) { %>
	go func() {
		
		baseServer,err := NewGRPCBaseServer(grpcListenPort,logger)
		if err !=nil {
			logger.Log("transport", "GRPC", "addr", grpcListenPort, "err", err.Error())
			errs <- err
		}
		grpcListener, err := net.Listen("tcp", grpcListenPort)
		logger.Log("transport", "GRPC", "addr", grpcListenPort)
		if err != nil {
			logger.Log("transport", "GRPC", "addr", grpcListenPort, "err", err.Error())
			errs <- err
		}

		grpcServer := NewGRPCServer(s)
		pb.RegisterBeKeysGoServiceServer(baseServer, grpcServer)
		reflection.Register(baseServer)

		err = baseServer.Serve(grpcListener)
		if err != nil {
			logger.Log("transport", "GRPC", "addr", grpcListenPort, "err", err.Error())
			errs <- err
		}

	}()
	<% } %>
	logger.Log("exit", <-errs)
}
