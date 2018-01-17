package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// Endpoints - Struct containing all endpoints
type Endpoints interface {
<% for(endpoint of endpoints) { %>
	<%= endpoint.methodName %>() endpoint.Endpoint
<% } %>
}

type endpoints struct {
	service BeKeysServiceInterface
}

// NewEndpoints - Creates new Endpoints
func NewEndpoints(service <%= serviceCamelCase %>Interface) Endpoints {
	return endpoints{service}
}
<% for(endpoint of endpoints) { %>
func (e endpoints) <%= endpoint.methodName %>() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{},error) {
		req,ok := request.(<%= endpoint.methodName %>Request)

		if !ok{
			return nil, NewError(500,"Failed to parse <%= endpoint.methodName %>Request")
		}
		// Endpoint logic
		response,err := e.service.<%= endpoint.methodName %>(ctx,req)
		return response, err
	}
}
<% } %>
