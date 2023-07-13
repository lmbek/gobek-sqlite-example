package html

import (
	"api/functions"
	"fmt"
	"net/http"
)

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	fmt.Println("1")
	err := functions.HandleAsLastEndpoint(endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println("2")
	switch request.Method {
	case http.MethodGet:
		return Get(response, request)
	case http.MethodPost:
		return nil, functions.NotSupportedError()
	case http.MethodPut:
		return nil, functions.NotSupportedError()
	case http.MethodDelete:
		return nil, functions.NotSupportedError()
	}

	return functions.InvalidEndpoint()
}
