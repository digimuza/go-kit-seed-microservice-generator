package main

import (
	_ "encoding/json"
)

// ServiceError - Extended error struct
type ServiceError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
}

// NewError - Creates new ServiceError
func NewError(statusCode int, err string) ServiceError {
	return ServiceError{statusCode, err}
}

func (err ServiceError) Error() string {
	return err.Err
}
