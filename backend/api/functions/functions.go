package functions

import (
	"errors"
	"fmt"
	"strings"
)

func CurrentLayer(endpoint string) string {
	// /api/personal//user/995/name

	splitted := strings.Split(endpoint, "/")
	// ["", "api", "user", "995", "name"]
	if len(splitted) >= 1 {
		return splitted[0] // /personal//user/995/name
	}
	return ""
}

func ReducedEndpoint(endpoint string) string {
	splitted := strings.Split(endpoint, "/")
	if len(splitted) >= 1 {
		return strings.Join(splitted[1:], "/")
	}
	return ""
}

func ReducedFirstLayerEndpoint(endpoint string) string {
	splitted := strings.Split(endpoint, "/")
	if len(splitted) >= 2 {
		return strings.Join(splitted[2:], "/")
	}
	return ""
}

func StandardEndpointData(endpoint string) (string, string) {
	reduced := ReducedEndpoint(endpoint)
	current := CurrentLayer(reduced)
	return reduced, current
}

// HandleAsLastEndpoint - This function is called when an endpoint should be handled as the last endpoint
// example:
// /api/users/something = invalid
// /api/users/ = [995]
func HandleAsLastEndpoint(endpoint string) error {
	reduced := ReducedEndpoint(endpoint)
	current := CurrentLayer(reduced)
	if current == "" {
		return nil
	}
	return InvalidEndpointError()
}

func HandleAsLastEndpointOLD(endpoint string) error {
	split := strings.Split(ReducedEndpoint(endpoint), "/") // example: name/something/something2
	fmt.Println("split")
	fmt.Println(split)
	if len(split) > 1 {
		// make sure that an endpoints such as {identifier}/name/ is not treated as {identifier}/name/s
		// if we have anything after our last endpoint: (/name/, /age/, etc.) then we return invalid endpoint
		if len(split[1]) > 0 {
			return InvalidEndpointError()
		}
	}
	return nil
}

func InvalidEndpoint() (any, error) {
	return nil, InvalidEndpointError()
}

func InvalidEndpointError() error {
	return errors.New("invalid endpoint")
}

func NotSupportedError() error {
	return errors.New("Sorry, the requested method is not supported for this path")
}
