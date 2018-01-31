package internal

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// Controller - all microservice available endpoints
type Controller interface {
	SampleEndPoint() Endpoint
}

// Endpoint - endpoint controler interface
type Endpoint interface {
	Server() *grpctransport.Server
	Endpoint(ctx context.Context, request interface{}) (interface{}, error)
	Decoder(ctx context.Context, request interface{}) (interface{}, error)
	Encoder(ctx context.Context, request interface{}) (interface{}, error)
}

func newController(service Service) Controller {
	return controller{service}
}

type controller struct {
	service Service
}

func newGRPCTransportServer(g Endpoint) *grpctransport.Server {
	server := grpctransport.NewServer(
		g.Endpoint,
		g.Decoder,
		g.Encoder,
	)
	return server
}

// SampleEndPoint

func (e controller) SampleEndPoint() Endpoint {
	return sampleEndPoint{e.service}
}

type sampleEndPoint struct {
	service Service
}

func (g sampleEndPoint) Server() *grpctransport.Server {
	return newGRPCTransportServer(g)
}

func (g sampleEndPoint) Decoder(ctx context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("SampleEndPoint endpoint is not implemented")
}

func (g sampleEndPoint) Endpoint(ctx context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("SampleEndPoint endpoint is not implemented")
}

func (g sampleEndPoint) Encoder(ctx context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("SampleEndPoint endpoint is not implemented")
}
