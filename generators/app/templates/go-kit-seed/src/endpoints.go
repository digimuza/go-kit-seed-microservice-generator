package main

import (
	"log"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// Endpoints - Struct containing all endpoints
type Endpoints struct {
<% for(endpoint of endpoints) { %>
	<%= endpoint.methodName %> endpoint.Endpoint
<% } %>
}
// NewEndpoints - Creates new Endpoints
func NewEndpoints(service <%= serviceCamelCase %>Interface) Endpoints {
	return Endpoints{
	<% for(endpoint of endpoints) { %>
		<%= endpoint.methodName %>: new<%= endpoint.methodName %>Endpoint(service),
	<% } %>
	}
}
<% for(endpoint of endpoints) { %>
func new<%= endpoint.methodName %>Endpoint(service <%= serviceCamelCase %>Interface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(<%= endpoint.methodName %>Request)
		
		// Endpoint logic

		response := <%= endpoint.methodName %>Response{}
		return response, nil
	}
	
}
<% } %>
