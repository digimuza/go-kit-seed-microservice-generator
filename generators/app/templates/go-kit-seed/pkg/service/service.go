package service

import (
	"context"
	"fmt"

	"github.com/asaskevich/govalidator"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

// Middleware - service middleware type
type Middleware func(next <%= serviceName %>) <%= serviceName %>

// <%= serviceName %> - service interface
type <%= serviceName %> interface {<% for (endpoint of methods) { %>
	<%= endpoint.methodName %>(ctx context.Context, req <%= endpoint.methodName %>Request) (<%= endpoint.methodName %>Response, error)<% } %>
}

// <%= serviceName %>Config - service configuration and dependeny injection object
type <%= serviceName %>Config struct{
	Logger log.Logger
	Created metrics.Counter
	Quered metrics.Counter
}

// NewService - Create BucketsService
func NewService(config <%= serviceName %>Config) <%= serviceName %> {
	return service{config}
}

type service struct{
	<%= serviceName %>Config
}

type (<% for (endpoint of methods) { %>
	// <%= endpoint.methodName %>Request - Request object
	<%= endpoint.methodName %>Request struct {
		Data string
	}
	// <%= endpoint.methodName %>Response - Response object
	<%= endpoint.methodName %>Response struct {
		Data string
	}<% } %>
)

<% for (endpoint of methods) { %>
func (s service) <%= endpoint.methodName %>(ctx context.Context, req <%= endpoint.methodName %>Request) (res <%= endpoint.methodName %>Response, err error) {
	err = errors.New("NOT IMPLEMENTED")
	return 
}<% } %>