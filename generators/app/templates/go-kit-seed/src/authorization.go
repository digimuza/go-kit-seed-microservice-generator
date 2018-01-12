package main

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

// JWTUserPayload with valid uuid
type JWTUserPayload struct {
	UID string `valid:"uuidv4, required"`
}

func getCurentUserUUID(ctx context.Context) (string, error) {
	payload, err := getCurrentUserPayload(ctx)

	if err != nil {
		return "", err
	}
	return payload.UID, nil
}

func getCurrentUserPayload(ctx context.Context) (JWTUserPayload, error) {
	meta, _ := metadata.FromIncomingContext(ctx)

	jwtTokkenArray := meta["authentication"]

	if len(jwtTokkenArray) == 0 {
		return JWTUserPayload{}, NewError(403, "Access Denied, failed to find authentication tokken in GRPC call metadata")
	}
	jwtTokken := jwtTokkenArray[0]

	return extractUserIDFromJWT(jwtTokken, "secret")
}

func extractUserIDFromJWT(tokenString string, signature string) (JWTUserPayload, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})

	if err != nil {
		return JWTUserPayload{}, err
	}

	if claims["uuid"] == nil {
		return JWTUserPayload{}, errors.New("No uuid property found in claims map")
	}

	payload := JWTUserPayload{
		UID: claims["uuid"].(string),
	}

	_, err = govalidator.ValidateStruct(payload)

	if err != nil {
		return JWTUserPayload{}, err
	}

	return payload, nil

}
