package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)
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
			userID, ok := vars["userID"]
			if !ok {
				return nil, NewError(http.StatusBadRequest, "UserID not found")
			}
			return obj, nil
		<% } %> 
	}
 <% } %>