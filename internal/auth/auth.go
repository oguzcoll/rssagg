package auth

import (
	"errors"
	"net/http"
	_ "net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an http request
// example:
//
//	Authorization: ApiKey {insert api key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no API Key provided")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid API Key provided")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header ")
	}
	return vals[1], nil
}
