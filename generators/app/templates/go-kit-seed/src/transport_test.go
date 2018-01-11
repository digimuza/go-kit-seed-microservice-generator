package main

import (
	"os"
	"context"
	"fmt"
	"testing"
	<% if(grpc){ %>
	"<%= org %>/<%= appName %>/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	<% } %> 
)
<% if(grpc){ %>
<% for(endpoint of endpoints) { %>
func Test<%= endpoint.methodName %>GRPCConnection(t *testing.T){
	client, conn, err := NewGRPCClient()

	if err != nil {
		t.Errorf(err.Error())
	}
	request := &pb.<%= endpoint.methodName %>Request{}

	response, callError := client.<%= endpoint.methodName %>(context.Background(), request)

	t.Errorf("Tests is not implemented")

	defer conn.Close()
}
<% } %>
<% } %> 