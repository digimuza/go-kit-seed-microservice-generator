package main

import (
	"context"
	<% if(http){ %>
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	
	"github.com/gorilla/mux"
	<% } %>
	<% if(grpc){ %>
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"<%= org %>/<%= appName %>/pkg/pb"
	<% } %>
)
<% if(http){ %>
// NewHTTPHandler - Creates new http.Handler
func NewHTTPHandler(s <%= serviceCamelCase %>Interface, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r := mux.NewRouter()
	e := NewEndpoints(s)

	<% for(endpoint of endpoints) { %>
	r.Methods("<%= endpoint.method %>").Path("<%= endpoint.path %>").Handler(httptransport.NewServer(
		e.<%= endpoint.methodName %>,
		decode<%= endpoint.methodName %>Request,
		encodeResponse,
		options...,
	))
	<% } %>
	return r
}
<% } %>
<% if(grpc){ %>

type grpcServer struct {
	<% for(endpoint of endpoints) { %>
	<%= endpoint.methodName.toLowerCase() %>    grpctransport.Handler
	<% } %>
	
}

func NewGRPCServer(s <%= serviceCamelCase %>Interface) pb.<%= serviceCamelCase %>Server {
	e := NewEndpoints(s)
	return &grpcServer{

		<% for(endpoint of endpoints) { %>
		<%= endpoint.methodName.toLowerCase() %>: grpctransport.NewServer(
			e.<%= endpoint.methodName %>,
			decodeGRPC<%= endpoint.methodName %>Request,
			encodeGRPC<%= endpoint.methodName %>Response,
		),
		<% } %>
	}
}

//NewGRPCBaseServer Generate Base server with certifications
func NewGRPCBaseServer(grpcServerTarget string, logger log.Logger) (*grpc.Server,error){
	// Read cert and key file

	creds, err := credentials.NewServerTLSFromFile(
		"/cert/backend.cert",
		"/cert/backend.key",
	)
	if err != nil {
		return nil,err
	}
	// Use Credentials in gRPC server options
	serverOption := grpc.Creds(creds)

	baseServer := grpc.NewServer(serverOption)

	return baseServer,nil
}

<% for(endpoint of endpoints) { %>
func (server *grpcServer) <%= endpoint.methodName %>(ctx context.Context, req *pb.<%= endpoint.methodName %>Request) (*pb.<%= endpoint.methodName %>Response, error) {
	_, rep, err := server.<%= endpoint.methodName.toLowerCase() %>.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.<%= endpoint.methodName %>Response), nil
}
<% } %>
<% } %>
