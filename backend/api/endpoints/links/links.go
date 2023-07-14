package links

import (
	"database/sqlite"
)

type Link struct {
	Id   int    "json:\"id\""
	Name string "json:\"name\""
}

func Get() (any, error) {
	database := new(sqlite.Database)

	err := database.Connect()
	defer database.Connection.Close()
	if err != nil {
		return nil, err
	}

	// if the table does not exist, create it (empty)
	err = database.CreateTable(sqlite.DataPath + "sql/links/create-table-links.sql")
	if err != nil {
		return nil, err
	}

	// check for data in the table
	result, err := database.SelectData(sqlite.DataPath + "sql/links/select-all-links.sql")
	if err != nil {
		return nil, err
	}

	var output []Link

	for result.Next() {
		link := new(Link)
		err := result.Scan(&link.Id, &link.Name)
		if err != nil {
			return nil, err
		}

		output = append(output, *link)
	}

	// Check if output is empty, return an empty slice instead of nil
	if len(output) == 0 {
		return []Link{}, nil
	}

	// if we have data, send data
	return output, nil
}
