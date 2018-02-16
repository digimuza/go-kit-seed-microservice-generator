package service

import (
	"encoding/json"
	"errors"

	pb "dev.adeoweb.biz/pas/be-buckets/pkg/pb/buckets"
	"github.com/gogo/protobuf/proto"
)

// ErrorType - Error maping to proto errors type
type ErrorType string

const (
	// ValidationError - validation error type
	ValidationError ErrorType = "type.googleapis.com/buckets.ValidationError"

	// JSONError - validation error type
	JSONError ErrorType = "type.googleapis.com/buckets.JSONError"

	// AccessDenied - permission error
	AccessDenied ErrorType = "type.googleapis.com/buckets.AccessDenied"
)

// ProtoType - returns proto struct
func (e ErrorType) ProtoType() (proto.Message, error) {
	errorTypes := map[ErrorType]proto.Message{
		ValidationError: &pb.ValidationError{},
	}

	t, ok := errorTypes[e]
	if !ok {
		return nil, errors.New("Failed to find proto type error definition")
	}

	return t, nil
}

// ErrorStatus -
type ErrorStatus interface {
	Message() string
	Details() []ErrorDetail
}

// ErrorDetail -
type ErrorDetail interface {
	Type() ErrorType
	Value() ([]byte, error)
	Error() string
}

type errorStatus struct {
	message string
	details []ErrorDetail
}

func (e errorStatus) Message() string {
	return e.message
}
func (e errorStatus) Details() []ErrorDetail {
	return e.details
}

// NewServiceError - Create service error
func NewServiceError(message string, details []ErrorDetail) ErrorStatus {
	return errorStatus{message, details}
}

// NewValidationErrorDetail - Create validation error
func NewValidationErrorDetail(message string) ErrorDetail {
	return ValidationErrorDetail{
		message,
	}
}

// ValidationErrorDetail - validation error
type ValidationErrorDetail struct {
	Message string
}

// Type - Validation error type
func (v ValidationErrorDetail) Type() ErrorType {
	return ValidationError
}

// Value - error value in bytes
func (v ValidationErrorDetail) Value() ([]byte, error) {
	return []byte(v.Message), nil
}

// Error - Provides error string
func (v ValidationErrorDetail) Error() string {
	return v.Message
}

// NewAccessDeniedErrorDetail - New access denied error
func NewAccessDeniedErrorDetail(message string) ErrorDetail {
	return AccessDeniedErrorDetail{
		Message: message,
	}
}

// AccessDeniedErrorDetail - access denied error
type AccessDeniedErrorDetail struct {
	Message string
}

// Type - Validation error type
func (v AccessDeniedErrorDetail) Type() ErrorType {
	return AccessDenied
}

// Value - error value in bytes
func (v AccessDeniedErrorDetail) Value() ([]byte, error) {
	return []byte(v.Message), nil
}

// Error - Provide string error represantation
func (v AccessDeniedErrorDetail) Error() string {
	return v.Message
}

// NewJSONErrorDetail - Flexible error
func NewJSONErrorDetail(i map[string]interface{}) ErrorDetail {
	return JSONErrorDetail{
		i,
	}
}

// JSONErrorDetail -  Flexible error
type JSONErrorDetail struct {
	Data map[string]interface{}
}

// Type - Validation error type
func (v JSONErrorDetail) Type() ErrorType {
	return JSONError
}

// Value - error value in bytes
func (v JSONErrorDetail) Value() ([]byte, error) {
	return json.Marshal(v.Data)
}

// Error - Provide string error represantation
func (v JSONErrorDetail) Error() string {
	err, _ := json.Marshal(v.Data)
	return string(err)
}
