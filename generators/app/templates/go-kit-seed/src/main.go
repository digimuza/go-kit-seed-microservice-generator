package main

import (
	"fmt"
	"net/http"
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
	var s <%= serviceCamelCase %>Interface
	{
		s = New<%= serviceCamelCase %>()
		s = NewLoggingMiddleware(s, logger)
	}
	var h http.Handler
	{
		h = NewHTTPHandler(s, logger)
	}
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", ":80")
		errs <- http.ListenAndServe(":80", h)
	}()

	logger.Log("exit", <-errs)
}
