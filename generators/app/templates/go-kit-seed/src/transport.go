package main

import (
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

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
