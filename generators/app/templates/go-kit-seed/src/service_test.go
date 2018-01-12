package main

import (
	"fmt"
	"testing"
)

<% for(endpoint of endpoints) { %>
func Test<%= endpoint.methodName %>(t *testing.T){

	service := New<%= serviceCamelCase %>()

	response, err := service.<%= endpoint.methodName %>(nil, <%= endpoint.methodName %>Request{})
	
	// Fail if any error
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(response)
	
}
<% } %>