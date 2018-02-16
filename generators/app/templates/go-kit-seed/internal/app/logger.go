package app

import (
	"os"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// NewLogger - create logger
func NewLogger(config utils.Configuration) log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, getLoggerOption(config))
		logger = level.NewInjector(logger, level.InfoValue())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	return logger
}

func getLoggerOption(config utils.Configuration) level.Option {
	if config.GetLogEnabled() != "true" {
		return level.AllowNone()
	}
	switch config.GetLogLevel() {
	case "debug":
		return level.AllowDebug()
	case "info":
		return level.AllowInfo()
	case "warn":
		return level.AllowWarn()
	case "error":
		return level.AllowError()
	default:
		return level.AllowAll()
	}
}
