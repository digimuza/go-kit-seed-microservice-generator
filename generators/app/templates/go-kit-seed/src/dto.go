package main

<% for(endpoint of endpoints) { %>

// <%= endpoint.methodName %>Request Data transfer object
type <%= endpoint.methodName %>Request struct {<% for(p of endpoint.parrams.request) { %>
	<%= p.goKey %> string `json:"<%= p.jsonKey %>"`
<% } %>
}

// <%= endpoint.methodName %>Response Data transfer object
type <%= endpoint.methodName %>Response struct {<% for(p of endpoint.parrams.response) { %>
	<%= p.goKey %> string `json:"<%= p.jsonKey %>"`
<% } %>
}

<% } %>