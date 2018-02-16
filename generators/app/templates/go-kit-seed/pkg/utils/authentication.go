package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
)

// JWTPayload - contains user information
type JWTPayload struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

// GetJWTPayload - returns jwt payload containing user information
func GetJWTPayload(ctx context.Context) (JWTPayload, error) {
	if ctx == nil {
		return JWTPayload{}, errors.New("Context can't be nil")
	}
	jwt, err := GetJWTToken(ctx)
	if err != nil {
		return JWTPayload{}, err
	}
	tokenParts := strings.Split(jwt, ".")
	if len(tokenParts) != 3 {
		return JWTPayload{}, errors.New("Malformed jwt token: failed to parse")
	}
	bytes, err := base64.RawStdEncoding.DecodeString(tokenParts[1])
	var payload JWTPayload
	json.Unmarshal(bytes, &payload)
	return payload, nil
}

// GetJWTToken - returns jwt token from metadata
func GetJWTToken(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errors.New("Context can't be nil")
	}
	meta, _ := metadata.FromIncomingContext(ctx)
	t := meta["authentication"]
	if len(t) == 0 {
		return "", errors.New("Access Denied, failed to find `authentication` token in gRPC call metadata")
	}
	return t[0], nil
}
