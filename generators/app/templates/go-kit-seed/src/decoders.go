package main

import (
	"context"
	<% if(http){ %>
	"encoding/json"
	"io/ioutil"
	"net/http"
	valid "github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	<% } %> 
	<% if(grpc){ %>
	"<%= app_name %>/pkg/pb"
	<% } %> 
)
<% if(http){ %>
<% for(endpoint of endpoints) { %>
func decode<%= endpoint.methodName %>Request(_ context.Context, r *http.Request) (interface{}, error) {
	var obj <%= endpoint.methodName %>Request
	<% if(endpoint.method === "POST" || endpoint.method === "PATCH"){ %>
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			return obj, err
		}

		err = json.Unmarshal(body, &obj)
		if err == nil {
			_, err = valid.ValidateStruct(obj)
		}
		if err != nil {
			err = NewError(http.StatusBadRequest, err.Error())
		}
		return obj, err
	<% } %> 
	<% if(endpoint.method === "GET" || endpoint.method === "DELETE"){ %>
		vars := mux.Vars(r)
		_, ok := vars["userID"]
		if !ok {
			return nil, NewError(http.StatusBadRequest, "UserID not found")
		}
		return obj, nil
	<% } %> 
	}
 <% } %>
 <% } %> 

 <% if(grpc){ %>
//GRPC Encoders - proto file should be builded and in
<% for(endpoint of endpoints) { %>
decodeGRPC<%= endpoint.methodName %>Request(_ context.Context, grpcRequest interface{}) (interface{}, error){
	req:= grpcRequest.(*pb.<%= endpoint.methodName %>Request)
	return <%= endpoint.methodName %>Request{response},nil
}
decodeGRPC<%= endpoint.methodName %>Response(_ context.Context, grpcResponse interface{}) (interface{}, error){
	response:= grpcResponse.(*pb.<%= endpoint.methodName %>Response)
	return <%= endpoint.methodName %>Response{response},nil
}
<% } %>
<% } %> 