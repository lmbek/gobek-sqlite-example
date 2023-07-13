package signout

import (
	"api/endpoints/auth/session"
	"net/http"
)

func SignOut(response http.ResponseWriter, request *http.Request) (any, error) {
	session.Close(response, request)
	return "Successfully signed out", nil
}
