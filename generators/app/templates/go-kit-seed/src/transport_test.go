package main

import (
	"context"
	"fmt"
	"testing"
	<% if(grpc){ %>
	"<%= org %>/<%= appName %>/pkg/pb"
	<% } %> 
)

var authentication = Authentication{
	JWT: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiOWRkOWZhZGYtM2Q5Ni00ZTA2LTk2YjEtOGEyYWZiMGJlNTQ4In0.wXuX4zKymPBFpKP4gKc4d56LCO2qjKScifVuttEn0Eo",
}


<% if(grpc){ %>
<% for(endpoint of endpoints) { %>
func Test<%= endpoint.methodName %>GRPCConnection(t *testing.T){
	client, conn, err := NewGRPCClient(authentication)

	if err != nil {
		t.Errorf(err.Error())
	}
	request := &pb.<%= endpoint.methodName %>Request{}

	response, callError := client.<%= endpoint.methodName %>(context.Background(), request)

	// Test response
	fmt.Println(response)

	if callError !=nil{
		t.Errorf("Tests is not implemented")
	}

	defer conn.Close()
}
<% } %>
<% } %> 