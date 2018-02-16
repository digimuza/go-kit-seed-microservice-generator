package endpoints

import (
	"context"
	"fmt"
	"reflect"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
)

// NewEndpoints - create struct that contains all endpoints
func NewEndpoints(s service.<%= serviceName %>, logger log.Logger, duration metrics.Histogram, tracer stdopentracing.Tracer) ControllerEndpoints {
	return endpoints{s, logger, duration, tracer}
}

// ControllerEndpoints -
type ControllerEndpoints interface {<% for (endpoint of methods) { %>
	<%= endpoint.methodName %>() endpoint.Endpoint
	<%= endpoint.methodName %>() endpoint.Endpoint<% } %>
}

type endpoints struct {
	service  service.BucketsService
	logger   log.Logger
	duration metrics.Histogram
	tracer   stdopentracing.Tracer
}


<% for (endpoint of methods) { %>
// <%= endpoint.methodName %> - endpoint
func (e endpoints) <%= endpoint.methodName %>() endpoint.Endpoint {
	var endpoint endpoint.Endpoint
	{
		endpoint = func(ctx context.Context, request interface{}) (interface{}, error) {
			req, ok := request.(service.<%= endpoint.methodName %>Request)
			if !ok {
				return nil, fmt.Errorf("incorect interface{} expected [service.<%= endpoint.methodName %>Request] got [%s]", reflect.TypeOf(request))
			}
			return e.service.<%= endpoint.methodName %>(ctx, req)
		}
		endpoint = e.endpointMiddlewares("<%= endpoint.methodName %>", endpoint)
	}

	return endpoint
}<% } %>

func (e endpoints) endpointMiddlewares(method string, en endpoint.Endpoint) endpoint.Endpoint {
	en = NewLoggingMiddleware(log.With(e.logger, "method", method))(en)
	en = opentracing.TraceServer(e.tracer, method)(en)
	en = NewInstrumentingMiddleware(e.duration.With("method", method))(en)
	return en
}
