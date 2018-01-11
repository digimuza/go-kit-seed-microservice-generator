package main

import (<% if(grpc){ %>
	"fmt"
	"os"
	"<%= org %>/<%= appName %>/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"<% } %> 
)

// NewGRPCClient - New service grpc client
func NewGRPCClient() (pb.<%= serviceCamelCase %>Client, error) {
	credsClient, err := credentials.NewClientTLSFromFile("../cert/frontend.cert", "")
	if err != nil {
		return nil, err
	}

	dialTarget := fmt.Sprintf("%s.%s:%d", os.Getenv("APP_NAME"), "passcamp.doc", 443)

	conn, err := grpc.Dial(dialTarget, grpc.WithTransportCredentials(credsClient))
	if err != nil {
		return nil, err
	}
	client := pb.New<%= serviceCamelCase %>Client(conn)

	return client, nil
}
