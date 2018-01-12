package main

import (<% if(grpc){ %>
	"fmt"
	"os"
	"<%= org %>/<%= appName %>/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"<% } %> 
)


// Authentication holds valid current user uuid
type Authentication struct {
	JWT string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authentication": a.JWT,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security. Implementing PerRPCCredentials interface
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

// NewGRPCClient - New service grpc client
func NewGRPCClient(auth Authentication) (pb.<%= serviceCamelCase %>Client,*grpc.ClientConn, error) {
	credsClient, err := credentials.NewClientTLSFromFile("../cert/frontend.cert", "")
	if err != nil {
		return nil, err
	}

	dialTarget := fmt.Sprintf("%s.%s:%d", os.Getenv("APP_NAME"), "passcamp.doc", 443)

	conn, err := grpc.Dial(
		dialTarget,
		grpc.WithTransportCredentials(credsClient),
		grpc.WithPerRPCCredentials(&auth),
	)
	if err != nil {
		return nil, err
	}
	client := pb.New<%= serviceCamelCase %>Client(conn)

	return client, nil
}
