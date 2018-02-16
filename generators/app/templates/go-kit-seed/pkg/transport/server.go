package transport

import (
	"context"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/endpoints"
	pb "dev.adeoweb.biz/pas/<%= appName %>/pkg/pb/buckets"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewGRPCServer - Generate Base server with certifications
func NewGRPCServer(config utils.Configuration) (*grpc.Server, error) {
	var baseServer *grpc.Server
	if config.IsSSLEnabled() {
		creds, err := credentials.NewServerTLSFromFile(
			config.GetSSLCertPath(),
			config.GetSSLKeyPath(),
		)
		if err != nil {
			return nil, err
		}
		serverOption := grpc.Creds(creds)
		baseServer = grpc.NewServer(serverOption)
	} else {
		baseServer = grpc.NewServer()
	}

	return baseServer, nil
}

type server struct {
<% for (endpoint of methods) { %>
	<%= endpoint.methodName %>   grpctransport.Handler
<% } %>
}

// NewServer - create proto file server
func NewServer(controller endpoints.Controller, tracer stdopentracing.Tracer, logger log.Logger) pb.BucketsServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &server{
		<% for (endpoint of methods) { %>
			<%= endpoint.methodName %>: grpctransport.NewServer(
				controller.Endpoints.<%= endpoint.methodName %>(),
				decode<%= endpoint.methodName %>Request,
				encode<%= endpoint.methodName %>Response,
				append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "<%= endpoint.methodName %>", logger)))...,
			),
		<% } %>
	}
}

<% for (endpoint of methods) { %>
func (s *server) <%= endpoint.methodName %>(ctx context.Context, req *pb.<%= endpoint.methodName %>Request) (*pb.<%= endpoint.methodName %>Response, error) {
	_, rep, err := s.<%= endpoint.methodName %>.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.<%= endpoint.methodName %>Response), nil
}
<% } %>	