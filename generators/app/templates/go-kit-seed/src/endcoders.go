package main

import (
	"context"
	<% if(http){ %>
	"encoding/json"
	"net/http"
	<% } %>
	<% if(grpc){ %>
	"<%= app_name %>/pkg/pb"
	<% } %>
)

<% if(http){ %>
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	sErr, ok := err.(ServiceError)
	if !ok {
		sErr = NewError(http.StatusInternalServerError, err.Error())
	}
	w.WriteHeader(sErr.StatusCode)
	json.NewEncoder(w).Encode(sErr)
}
<% } %>

<% if(grpc){ %>
	
//GRPC Encoders - proto file should be builded and in
<% for(endpoint of endpoints) { %>
encodeGRPC<%= endpoint.methodName %>Request(_ context.Context, request interface{}) (interface{}, error){
	req:= request.(<%= endpoint.methodName %>Request)
	return &pb.<%= endpoint.methodName %>Request{},nil
}
encodeGRPC<%= endpoint.methodName %>Response(_ context.Context, request interface{}) (interface{}, error){
	req:= request.(<%= endpoint.methodName %>Response)
	return &pb.<%= endpoint.methodName %>Response{},nil
}
<% } %>

<% } %> 
