package main

import (
	"awpc/<%= appName %>/internal"
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		serviceConfig := internal.ServiceConfig{
			DB: createDBConnection(),
		}
		internal.ServeGRPC(serviceConfig, logger, errs)
	}()

	logger.Log("exit", <-errs)
}
