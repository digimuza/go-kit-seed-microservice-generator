package main

<% for(endpoint of endpoints) { %>

type <%= endpoint.methodName %>Request struct {}
type <%= endpoint.methodName %>Response struct {}

<% } %>