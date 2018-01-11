package main

<% for(endpoint of endpoints) { %>

// <%= endpoint.methodName %>Request Data transfer object
type <%= endpoint.methodName %>Request struct {}

// <%= endpoint.methodName %>Response Data transfer object
type <%= endpoint.methodName %>Response struct {}

<% } %>