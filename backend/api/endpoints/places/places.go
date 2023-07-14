package places

import (
	"database/sqlite"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Place struct {
	Id   int    "json:\"id\""
	Name string "json:\"name\""
}

func Get() (any, error) {
	database := sqlite.Database{}

	err := database.Connect()
	defer database.Connection.Close()
	if err != nil {
		return nil, err
	}

	// if the table does not exist, create it (empty)
	err = database.CreateTable(sqlite.DataPath + "sql/places/create-table-places.sql")
	if err != nil {
		return nil, err
	}

	// check for data in the table
	result, err := database.SelectData(sqlite.DataPath + "sql/places/select-all-places.sql")
	if err != nil {
		return nil, err
	}

	var output []Place

	for result.Next() {
		place := new(Place)
		err := result.Scan(&place.Id, &place.Name)
		if err != nil {
			return nil, err
		}

		output = append(output, *place)
	}

	// Check if output is empty, return an empty slice instead of nil
	if len(output) == 0 {
		return []Place{}, nil
	}

	// if we have data, send data
	myData := output
	return myData, nil
}

func Post(response http.ResponseWriter, request *http.Request) (any, error) {
	// If we got it as form data
	if request.Header.Get("content-type") == "application/x-www-form-urlencoded" {
		return PostWithForm(response, request)
	}

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

	// json data expected to be formed as struct Place
	place := Place{}

	// convert json to type
	err = json.Unmarshal([]byte(body), &place)
	if err != nil {
		return nil, err
	}

	return databasePost(place)
}

func PostWithForm(response http.ResponseWriter, request *http.Request) (any, error) {
	err := request.ParseForm()
	if err != nil {
		return nil, err
	}

	place := new(Place)

	place.Name = request.Form.Get("name")

	return databasePost(*place)
}

func databasePost(place Place) (any, error) {
	// connect to database
	database := sqlite.Database{}
	err := database.Connect()
	defer database.Connection.Close()
	if err != nil {
		return nil, err
	}

	// create database table if not exist
	err = database.CreateTable(sqlite.DataPath + "sql/places/create-table-places.sql")
	if err != nil {
		return nil, err
	}

	// insert into database
	err = database.InsertData(sqlite.DataPath+"sql/places/insert-into-places.sql", nil, place.Name)
	if err != nil {
		return nil, err
	}

	return "you have added data", nil
}
