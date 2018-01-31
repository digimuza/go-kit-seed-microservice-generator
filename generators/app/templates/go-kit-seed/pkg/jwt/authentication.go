package jwt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc/metadata"
)

// Authentication - data entity
type Authentication struct {
	UserID  string `valid:"uuidv4, required"`
	Email   string `valid:"email"`
	Payload map[string]interface{}
	JWT     string
}

// Auth - authenticates context object
func Auth(ctx context.Context) (Authentication, error) {
	meta, _ := metadata.FromIncomingContext(ctx)

	t := meta["authentication"]

	if len(t) == 0 {
		return Authentication{}, errors.New("Access Denied, failed to find Authentication token in gRPC call metadata")
	}
	jwt := t[0]

	payload, err := extractPayLoadFromJWT(jwt)

	if err != nil {
		return Authentication{}, err
	}

	userID, userIDExsist := payload["userId"]
	email, emailExsist := payload["email"]
	if !userIDExsist || !emailExsist {
		return Authentication{}, errors.New("incorrect payload data")
	}

	response := Authentication{
		UserID:  userID.(string),
		Email:   email.(string),
		JWT:     jwt,
		Payload: payload,
	}

	_, err = govalidator.ValidateStruct(response)

	if err != nil {
		return Authentication{}, err
	}

	return response, err
}

func extractPayLoadFromJWT(jwt string) (map[string]interface{}, error) {

	payload := make(map[string]interface{})

	// tokenParts[0] - jwt token base64UrlEncode(header)
	// tokenParts[1] - jwt token base64(payload)
	// tokenParts[2] - jwt token base64(sumCheck)
	tokenParts := strings.Split(jwt, ".")

	if len(tokenParts) != 3 {
		return nil, errors.New("Malformed token: failed to parse")
	}

	payloadBites, err := base64.StdEncoding.DecodeString(tokenParts[1])

	err = json.Unmarshal(payloadBites, &payload)

	if err != nil {
		return nil, err
	}

	return payload, err
}
