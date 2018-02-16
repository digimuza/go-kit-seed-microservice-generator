package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// NewLoggingMiddleware - Creates logging middleware for endpoint
func NewLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Debug(logger).Log("msg", "calling endpoint")
			defer level.Debug(logger).Log("msg", "called endoint")
			return next(ctx, request)
		}
	}
}
