package id

import (
	"api/endpoints/profile/id/age"
	"api/endpoints/profile/id/name"
	"api/functions"
	"errors"
	"net/http"
)

func HandleIdentifierEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	// get current endpoint (identifier)
	reducedIdentifierEndpoint, identifier := functions.StandardEndpointData(endpoint)

	// Check if identifier is empty
	if len(identifier) == 0 {
		return nil, errors.New("need identifier: /api/profile/{id}")
	}

	// if identifier is not empty we get the endpoints for identifier by reducing it further
	// get endpoints after identifier: {identifier}/name etc.
	reducedEndpoint, currentEndpoint := functions.StandardEndpointData(reducedIdentifierEndpoint)

	if len(currentEndpoint) == 0 { // if we could find user but there was not requested any endpoints after identifier: (example) 995
		userID, err := Get(identifier) // Search for user in database
		if err != nil {
			return nil, err // user do not exist or database could not connect, etc...
		}
		return userID, nil
	} else {
		switch currentEndpoint { // path: {id}/name/something
		case "name":
			return name.HandleEndpoint(reducedEndpoint, identifier, response, request)
		case "age":
			return age.HandleEndpoint(reducedEndpoint, identifier, response, request)
		}
	}
	return nil, functions.InvalidEndpointError()
}
