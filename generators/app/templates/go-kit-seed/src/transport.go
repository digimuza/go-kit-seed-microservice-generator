package main

import (

	"github.com/go-kit/kit/log"
	<% if(http){ %>
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	<% } %>
	
	<% if(grpc){ %>
	"<%= appName %>/pkg/pb"
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

func NewGRPCServer(s <%= serviceCamelCase %>Interface, logger log.Logger) pb.<%= serviceCamelCase %>Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	e := NewEndpoints(s)
	return &grpcServer{

		<% for(endpoint of endpoints) { %>
		<%= endpoint.methodName.toLowerCase() %>: grpctransport.NewServer(
			e.<%= endpoint.methodName %>,
			decodeGRPC<%= endpoint.methodName %>Request,
			encodeGRPC<%= endpoint.methodName %>Response,
			options...,
		),
		<% } %>
	}
}

<% for(endpoint of endpoints) { %>
func (server *grpcServer) Sum(ctx oldcontext.Context, req *pb.<%= endpoint.methodName %>Request) (*pb.<%= endpoint.methodName %>Response, error) {
	_, rep, err := server.<%= endpoint.methodName.toLowerCase() %>.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.<%= endpoint.methodName %>Response), nil
}
<% } %>


<% } %>