package main

import (
	"testing"
)

func TestNew<%= serviceCamelCase %>(t *testing.T){
	t.Errorf("Test is not implemented")
}
<% for(endpoint of endpoints) { %>
func Test<%= endpoint.methodName %>(t *testing.T){
	t.Errorf("Test is not implemented")
}
<% } %>