package png

import (
	"api/functions"
	"net/http"
)

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	err := functions.HandleAsLastEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

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
