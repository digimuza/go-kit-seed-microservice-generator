package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
)



// <%= serviceCamelCase %>Interface - Main service interface
type <%= serviceCamelCase %>Interface interface {
	<% for(endpoint of endpoints) { %>
		<%= endpoint.methodName %>(ctx context.Context, req <%= endpoint.methodName %>Request) (<%= endpoint.methodName %>Response, error)
	<% } %>
}

// <%= serviceCamelCase %> - Main service struct this struct should contain only bussiness logic
type <%= serviceCamelCase %> struct {}

// New<%= serviceCamelCase %> - Creates new <%= serviceCamelCase %>
func New<%= serviceCamelCase %>() <%= serviceCamelCase %>Interface {
	return <%= serviceCamelCase %>{}
}
<% for(endpoint of endpoints) { %>
func (service <%= serviceCamelCase %>) <%= endpoint.methodName %>(ctx context.Context, req <%= endpoint.methodName %>Request) (<%= endpoint.methodName %>Response, error){
	return <%= endpoint.methodName %>Response, nil
}
<% } %>
func getDBUrl() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
}