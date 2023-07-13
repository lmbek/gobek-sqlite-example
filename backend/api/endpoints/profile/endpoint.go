package profile

import (
	"api/endpoints/profile/id"
	"api/functions"
	"net/http"
)

// HandleEndpoint - If no endpoints, then we try to return all data on user (if user id exists),
// if they don't have a session, some or all parts of the api might not be returned
// if no user id with that identifier, then an error with "no such user" will be given
func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	currentEndpoint := functions.CurrentLayer(endpoint)

	// This is overwriting current endpoint, such as if we want some other features extra...
	switch currentEndpoint {
	// here we can add anything that should overwrite the identifier
	}

	// handle the rest as identifier
	return id.HandleIdentifierEndpoint(endpoint, response, request)
}
