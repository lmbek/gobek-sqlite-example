package api

import (
	"api/endpoints/auth"
	"api/endpoints/download"
	"api/endpoints/example"
	"api/endpoints/html"
	"api/endpoints/links"
	"api/endpoints/native"
	"api/endpoints/places"
	"api/endpoints/png"
	"api/endpoints/profile"
	"api/endpoints/profiles"
	"api/functions"
	"errors"
	"net/http"
)

// Example data (path): /api/user/995/name/something
// HandleEndpoint - This method can be mimiced, it reduces the endpoint and selects the current for use
func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	reducedEndpoint := functions.ReducedFirstLayerEndpoint(endpoint) // Example data: user/995/name/something
	currentEndpoint := functions.CurrentLayer(reducedEndpoint)       // Example data: user

	switch currentEndpoint {
	case "auth":
		return auth.HandleEndpoint(reducedEndpoint, response, request)
	case "download", "png", "html":
		return ModifyResponse(reducedEndpoint, response, request)
	case "example":
		return example.HandleEndpoint(reducedEndpoint, response, request)
	case "links":
		return links.HandleEndpoint(reducedEndpoint, response, request)
	case "native":
		return native.HandleEndpoint(reducedEndpoint, response, request)
	case "places":
		return places.HandleEndpoint(reducedEndpoint, response, request)
	case "profile":
		return profile.HandleEndpoint(reducedEndpoint, response, request)
	case "profiles":
		return profiles.HandleEndpoint(reducedEndpoint, response, request)
	}

	return nil, errors.New("invalid endpoint")
}

func ModifyResponse(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	current := functions.CurrentLayer(endpoint)
	switch current {
	case "download":
		return download.HandleEndpoint(endpoint, response, request)
	case "png":
		return png.HandleEndpoint(endpoint, response, request)
	case "html":
		return html.HandleEndpoint(endpoint, response, request)
	}

	return nil, errors.New("invalid endpoint")
}
