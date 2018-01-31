package internal

import (
	"context"
	"fmt"
	"net"
	"os"
	"awpc/<%= appName %>/pkg/pb"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// ServeGRPC - function that creates grpc server
func ServeGRPC(config ServiceConfig, logger log.Logger, errs chan error) {
	address := fmt.Sprintf("%s.%s:%s", os.Getenv("APP_NAME"), os.Getenv("DOMAIN"), os.Getenv("APP_PORT"))
	baseServer, err := NewGRPCBaseServer()
	if err != nil {
		logger.Log("transport", "GRPC", "err", err.Error())
		errs <- err
	}

	grpcListener, err := net.Listen("tcp", address)
	logger.Log("transport", "GRPC", "addr", address)
	if err != nil {
		logger.Log("transport", "GRPC", "err", err.Error())
		errs <- err
	}
	s := NewService(config)

	protoServer := NewProtoServer(s)
	pb.Register<%= serviceName %>Server(baseServer, protoServer)
	reflection.Register(baseServer)

	err = baseServer.Serve(grpcListener)
	if err != nil {
		logger.Log("transport", "GRPC", "err", err.Error())
		errs <- err
	}

}

// NewGRPCBaseServer - Generate Base server with certifications
func NewGRPCBaseServer() (*grpc.Server, error) {
	// Read cert and key file
	creds, err := credentials.NewServerTLSFromFile(
		os.Getenv("CERT_PATH"),
		os.Getenv("KEY_PATH"),
	)

	if err != nil {
		return nil, err
	}

	serverOption := grpc.Creds(creds)
	baseServer := grpc.NewServer(serverOption)

	return baseServer, nil
}

// NewBucketsServiceServer - create proto file server
func NewProtoServer(s Service) pb.<%= serviceName %>Server {
	controller := newController(s)
	return &transport{controller}
}

type transport struct {
	controler Controller
}

func (server *transport) SampleEndPoint(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	_, rep, err := server.controler.SampleEndPoint().Server().ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Data), nil
}

