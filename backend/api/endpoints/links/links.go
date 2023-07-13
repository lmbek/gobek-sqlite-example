package links

import (
	"database/sqlite"
	"errors"
)

func Get() (any, error) {
	database := sqlite.Database{}

	err := database.Connect()
	defer database.Connection.Close()
	if err != nil {
		return nil, err
	}

	err = database.CreateTable(sqlite.DataPath + "sql/links/create-table-links.sql")
	if err != nil {
		return nil, err
	}

	// if we have data, send data
	if true {
		myData := []string{"https://google.dk", "https://facebook.com", "https://linkedin.com"}
		return myData, nil
	}

	// else send the error (modify current error)
	return nil, errors.New("database could not connect")
}
