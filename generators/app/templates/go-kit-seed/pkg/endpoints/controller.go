package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/service"
	stdopentracing "github.com/opentracing/opentracing-go"
)

// Controller - contains all endpoints and middlawares
type Controller struct {
	Endpoints ControllerEndpoints
}

// NewController - creates a new controller
func NewController(s service.<%= serviceName %>, logger log.Logger, duration metrics.Histogram, tracer stdopentracing.Tracer) Controller {
	endpoints := NewEndpoints(s, logger, duration, tracer)
	return Controller{endpoints}
}

<% for (endpoint of methods) { %>
// <%= endpoint.methodName %> - endpoint
func (s Controller) <%= endpoint.methodName %>(ctx context.Context, req service.<%= endpoint.methodName %>Request) (res service.<%= endpoint.methodName %>Response, err error) {
	r, err := s.Endpoints.<%= endpoint.methodName %>()(ctx, req)
	if err != nil {
		return
	}

	res, ok := r.(service.<%= endpoint.methodName %>Response)
	if !ok {
		err = errors.New("Invalid response returned by Controller.<%= endpoint.methodName %>")
	}
	return
}
<% } %>