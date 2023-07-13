package auth

import (
	"api/endpoints/auth/active"
	"api/endpoints/auth/signin"
	"api/endpoints/auth/signout"
	"api/functions"
	"errors"
	"net/http"
)

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	_, current := functions.StandardEndpointData(endpoint)

	switch current {
	case "active":
		return active.HandleEndpoint(current, response, request)
	case "signin":
		return signin.HandleEndpoint(current, response, request)
	case "signout":
		return signout.HandleEndpoint(current, response, request)
	case "signup":
		return nil, errors.New("not implemented")
	}

	return nil, errors.New("You can use the following: /api/auth/active, /api/auth/signin, /api/auth/signout/, /api/auth/signup")
}
