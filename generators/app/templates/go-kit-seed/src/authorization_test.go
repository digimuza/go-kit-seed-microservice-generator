package main

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestExtractUserIDFromJWTValid(t *testing.T) {

	jwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiOWRkOWZhZGYtM2Q5Ni00ZTA2LTk2YjEtOGEyYWZiMGJlNTQ4In0.wXuX4zKymPBFpKP4gKc4d56LCO2qjKScifVuttEn0Eo"
	jwtSecret := "secret"

	payload, err := extractUserIDFromJWT(jwtString, jwtSecret)

	if err != nil {
		t.Errorf(err.Error())
	}

	if payload.UID != "9dd9fadf-3d96-4e06-96b1-8a2afb0be548" {
		t.Errorf("Something unexpected happaned. Decode payload ID is not == expected ID")
	}

}

func TestExtractUserIDFromJWTExpectToReturnError(t *testing.T) {

	// Invalid jwtString
	jwtString := "asd.eyJ1dWlkIjoiOasdWRkOWZhZGYtM2Q5Ni00ZTA2LTk2YjEtOGEyYWZiMGJlNTQ4In0.wXuX4zKymPBFpKP4gKc4d56LCO2qjKScifVuttEn0Eo"
	jwtSecret := "secret"

	_, err := extractUserIDFromJWT(jwtString, jwtSecret)

	if err == nil {
		t.Errorf(err.Error())
	}

	// No uid property in jwtString
	jwtString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkYyI6IjlkZDlmYWRmLTNkOTYtNGUwNi05NmIxLThhMmFmYjBiZTU0OCJ9.rfXJPHMRXuw8ZQi-y0CUYV4YC2OA850AtQh6V7uEiBU"

	_, err = extractUserIDFromJWT(jwtString, jwtSecret)

	if err == nil {
		t.Errorf(err.Error())
	}

}

func TestGetUserPayloadFromCtx(t *testing.T) {

	md := metadata.Pairs(
		"authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiOWRkOWZhZGYtM2Q5Ni00ZTA2LTk2YjEtOGEyYWZiMGJlNTQ4In0.wXuX4zKymPBFpKP4gKc4d56LCO2qjKScifVuttEn0Eo",
	)

	// create a new context with this metadata
	ctx := metadata.NewIncomingContext(context.Background(), md)

	payload, err := getCurrentUserPayload(ctx)

	if err != nil {
		t.Errorf(err.Error())
	}

	if payload.UID != "9dd9fadf-3d96-4e06-96b1-8a2afb0be548" {
		t.Errorf("Something unexpected happaned. Decode payload ID is not == expected ID")
	}

}

func TestFailGetUserPayloadFromCtx(t *testing.T) {

	md := metadata.Pairs(
		"authentication", "asdasd.asda.d",
	)

	// create a new context with this metadata
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := getCurrentUserPayload(ctx)

	if err == nil {
		t.Errorf("Expected error")
	}

	md = metadata.Pairs(
		"authentication", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkYyI6IjlkZDlmYWRmLTNkOTYtNGUwNi05NmIxLThhMmFmYjBiZTU0OCJ9.rfXJPHMRXuw8ZQi-y0CUYV4YC2OA850AtQh6V7uEiBU",
	)

	// create a new context with this metadata
	ctx = metadata.NewIncomingContext(context.Background(), md)

	_, err = getCurrentUserPayload(ctx)

	if err == nil {
		t.Errorf("Expected error didint provide uuid property in jwt token payload")
	}

}
