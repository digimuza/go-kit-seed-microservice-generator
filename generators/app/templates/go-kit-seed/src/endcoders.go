package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	sErr, ok := err.(ServiceError)
	if !ok {
		sErr = NewError(http.StatusInternalServerError, err.Error())
	}
	w.WriteHeader(sErr.StatusCode)
	json.NewEncoder(w).Encode(sErr)
}
