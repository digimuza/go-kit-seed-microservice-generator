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

func TestGRPCConnection(t *testing.T) {

	// Create credentials
	credsClient, err := credentials.NewClientTLSFromFile("../cert/frontend.cert", "")
	if err !=nil{
		t.Errorf(err.Error())
	}

	dialTarget := fmt.Sprintf("%s.%s:%d",os.Getenv("APP_NAME"),"passcamp.doc",443)

	conn, err := grpc.Dial(dialTarget, grpc.WithTransportCredentials(credsClient))
	if err != nil {
		t.Errorf(err.Error())
	} else {
	}
	client := pb.NewBeKeysGoServiceClient(conn)
	request := &pb.GetUserByIDRequest{Id: "SampleID"}

	_, err = client.GetUserByID(context.Background(), request)

	if err != nil {
		t.Errorf(err.Error())
	} 
	defer conn.Close()
}
