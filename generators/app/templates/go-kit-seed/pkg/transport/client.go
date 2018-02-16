package transport

import (
	"context"
	"errors"
	"fmt"
	"time"

	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/endpoints"
	pb "dev.adeoweb.biz/pas/<%= appName %>/pkg/pb/buckets"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/service"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"
)

// Authentication holds valid current user uuid
type Authentication struct {
	JWT string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authentication": a.JWT,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security. Implementing PerRPCCredentials interface
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

// NewClientConnection - creates new grpc connection
func NewClientConnection(jwt string, config utils.Configuration) (*grpc.ClientConn, error) {
	auth := Authentication{
		JWT: jwt,
	}
	dialTarget := fmt.Sprintf("%s:%s", "<%= appName %>", "443")
	if config.IsSSLEnabled() {
		credsClient, err := credentials.NewClientTLSFromFile(config.GetSSLCertPath(), "")
		if err != nil {
			return nil, err
		}
		conn, err := grpc.Dial(
			dialTarget,
			grpc.WithTransportCredentials(credsClient),
			grpc.WithPerRPCCredentials(&auth),
		)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}

	conn, err := grpc.Dial(
		dialTarget,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(&auth),
	)
	if err != nil {
		return nil, err
	}

	if conn == nil {
		return nil, errors.New("Failed establish connection")
	}

	return conn, nil
}

// NewEndpoints - NewClient endpoints
func NewEndpoints(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) endpoints.ControllerEndpoints {
	return clientEndpoints{
		conn,
		tracer,
		logger,
	}
}

type clientEndpoints struct {
	conn   *grpc.ClientConn
	tracer stdopentracing.Tracer
	logger log.Logger
}
<% for (endpoint of methods) { %>
	func (c clientEndpoints) <%= endpoint.methodName %>() endpoint.Endpoint {
		method := "<%= endpoint.methodName %>"
		limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
		var endpoint endpoint.Endpoint
		{
			endpoint = grpctransport.NewClient(
				c.conn,
				"buckets.<%= serviceName %>",
				method,
				encode<%= endpoint.methodName %>Request,
				decode<%= endpoint.methodName %>Response,
				pb.<%= endpoint.methodName %>Response{},
				grpctransport.ClientBefore(opentracing.ContextToGRPC(c.tracer, c.logger)),
			).Endpoint()
			endpoint = opentracing.TraceClient(c.tracer, method)(endpoint)
			endpoint = limiter(endpoint)
			endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
				Name:    method,
				Timeout: 30 * time.Second,
			}))(endpoint)
		}
	
		return endpoint
	}
<% } %>

// NewClient - create new entity event client
func NewClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) service.<%= serviceName %> {
	clientEndpoints := NewEndpoints(conn, tracer, logger)
	return endpoints.Controller{
		Endpoints: clientEndpoints,
	}
}
