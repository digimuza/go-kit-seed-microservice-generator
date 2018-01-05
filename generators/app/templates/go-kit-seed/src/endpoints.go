package main

import (
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
	return func(ctx context.Context, request interface{}) (interface{},error) {
		req,ok := request.(<%= endpoint.methodName %>Request)

		if !ok{
			return nil, NewError(500,"Failed to parse <%= endpoint.methodName %>Request")
		}
		// Endpoint logic
		response,err := service.<%= endpoint.methodName %>(ctx,req)
		return response, err
	}
	
}
<% } %>
