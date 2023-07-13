package signin

import (
	"api/endpoints/auth/session"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ConnectToDatabase() {
	if true {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}
}

// SignIn - In order to sign in, the following needs to happen:
// 1) Get Payload data from request
// 2) Check Payload data (username, password)
// 3) Give error if login credentials is wrong
// 4) Give Session Cookie if login credentials match
func SignIn(response http.ResponseWriter, request *http.Request) (any, error) {
	ConnectToDatabase()

	// If we got it as form data
	if request.Header.Get("content-type") == "application/x-www-form-urlencoded" {
		return SignInWithForm(response, request)
	}

	// Otherwise we expect to get it as json

	// Check if there is no request body (data)
	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}

	// Read the payload from the request body
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return nil, errors.New("Could not read request body (error): " + err.Error())
	}

	// Try to convert to json
	validJSON := json.Valid(body)
	if !validJSON {
		return nil, errors.New("Could not convert json to objects: json might be invalid")
	}

	// json data expected to be formed as struct Login
	type Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	login := Login{}

	// convert json to type
	err = json.Unmarshal([]byte(body), &login)
	if err != nil {
		return nil, err
	}

	if login.Username == "lars" && login.Password == "kage" {
		_, err = session.Create(response, request, "995")
		if err != nil {
			return nil, err
		}

		return "successfully signed in", nil
	}

	return nil, errors.New("login credentials is not correct")
}

func SignInWithForm(response http.ResponseWriter, request *http.Request) (any, error) {
	err := request.ParseForm()
	if err != nil {
		return nil, err
	}

	username := request.Form.Get("username")
	password := request.Form.Get("password")

	if username == "lars" && password == "kage" {
		_, err = session.Create(response, request, "995")
		if err != nil {
			return nil, err
		}

		return "successfully signed in", nil
	}

	return nil, errors.New("incorrect login credentials")
}
