package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("authorization header is not set") 
	}	

	apiKeys := strings.Split(value, " ")
	if len(apiKeys) != 2 {
		return "", errors.New("malformed auth header")
	}

	if apiKeys[0] != "Apikey" {
		return "", errors.New("malformed first part of auth header")
	}

	return apiKeys[1], nil
}