package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/ptypes"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/models"
	pb "dev.adeoweb.biz/pas/<%= appName %>/pkg/pb/buckets"
	"dev.adeoweb.biz/pas/<%= appName %>/pkg/service"
)
<% for (endpoint of methods) { %>
	func decode<%= endpoint.methodName %>Request(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.<%= endpoint.methodName %>Request)
		if !ok {
			return nil, fmt.Errorf("decode<%= endpoint.methodName %>Request(): incorect interface{} expected [*pb.<%= endpoint.methodName %>Request] got [%s]", reflect.TypeOf(request))
		}
	
		return service.<%= endpoint.methodName %>Request{}, nil
	}
<% } %>

<% for (endpoint of methods) { %>
	func decode<%= endpoint.methodName %>Response(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.<%= endpoint.methodName %>Response)
		response := service.<%= endpoint.methodName %>Response{}
		if !ok {
			return nil, fmt.Errorf("decode<%= endpoint.methodName %>Response(): incorect interface{} expected [*pb.<%= endpoint.methodName %>Response] got [%s]", reflect.TypeOf(request))
		}
		return response, nil
	}
<% } %>