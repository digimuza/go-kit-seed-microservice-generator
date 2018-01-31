package client

import (
	"awpc/be-buckets/pkg/pb"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Data - data set to initialize client
type Data struct {
	Target        string
	JWT           string
	CredetialPath string
}

// GetCrediantials create credentials
func (a *Data) GetCrediantials() (credentials.TransportCredentials, error) {
	cred, err := credentials.NewClientTLSFromFile(a.CredetialPath, "")
	if err != nil {
		return nil, err
	}

	return cred, nil
}

// GetRequestMetadata gets the current request metadata
func (a *Data) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authentication": a.JWT,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security. Implementing PerRPCCredentials interface
func (a *Data) RequireTransportSecurity() bool {
	return false
}

// NewGRPCClient - New service grpc client
func NewGRPCClient(clientData Data) (pb.BeBucketsServiceClient, *grpc.ClientConn, error) {

	cred, err := clientData.GetCrediantials()
	if err != nil {
		return nil, nil, err
	}

	conn, err := grpc.Dial(
		clientData.Target,
		grpc.WithTransportCredentials(cred),
		grpc.WithPerRPCCredentials(&clientData),
	)
	if err != nil {
		return nil, conn, err
	}

	client := pb.New<%= serviceCamelCase %>Client(conn)

	return client, conn, nil
}
