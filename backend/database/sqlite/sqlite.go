package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var DataPath = "./data/"

type Database struct {
	Connection *sql.DB
}

// Open a connection to the SQLite database.
func (database *Database) Connect() error {
	var err error
	database.Connection, err = sql.Open("sqlite3", DataPath+"gobek-example.db")
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) prepareStatement(SQLFile string) (*sql.Stmt, error) {
	query, err := os.ReadFile(SQLFile)
	if err != nil {
		return nil, err
	}

	statement, err := database.Connection.Prepare(string(query))
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (database *Database) CreateTable(SQLFile string) error {
	statement, err := database.prepareStatement(SQLFile)
	defer statement.Close()
	if err != nil {
		fmt.Println("aaa")
		fmt.Println(err)
		fmt.Println("bbb")
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) InsertData(SQLFile string, args ...any) error {
	statement, err := database.prepareStatement(SQLFile)
	defer statement.Close()
	if err != nil {
		return err
	}

	// Insert a new row into the "user" table using the prepared statement.
	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) SelectData(SQLFile string, args ...any) (*sql.Rows, error) {
	statement, err := database.prepareStatement(SQLFile)
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	// Modify the users data using the prepared statement.
	result, err := statement.Query(args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (database *Database) UpdateData(SQLFile string, args ...any) error {
	statement, err := database.prepareStatement(SQLFile)
	defer statement.Close()
	if err != nil {
		return err
	}

	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) DeleteData(SQLFile string, args ...any) error {
	statement, err := database.prepareStatement(SQLFile)
	defer statement.Close()
	if err != nil {
		return err
	}

	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func printAllUsersExample(allUsersData *sql.Rows) error {
	var id, age int
	var name, website string
	fmt.Print("id")
	fmt.Print(" ")
	fmt.Print("name")
	fmt.Print(" ")
	fmt.Print("age")
	fmt.Print(" ")
	fmt.Print("website")
	fmt.Println()
	for allUsersData.Next() { //while loop
		err := allUsersData.Scan(&id, &name, &age, &website)
		if err != nil {
			return err
		}
		fmt.Print(id)
		fmt.Print(" ")
		fmt.Print(name)
		fmt.Print(" ")
		fmt.Print(age)
		fmt.Print(" ")
		fmt.Print(website)
		fmt.Println()
	}

	return nil
}

func runExamples(database *Database) error {
	err := database.CreateTable(DataPath + "sql/examples/create-table-users.sql")
	if err != nil {
		return err
	}

	err = database.InsertData(DataPath+"sql/examples/insert-into-users.sql", nil, "lars3", 28, "")
	if err != nil {
		return err
	}

	result, err := database.SelectData(DataPath + "sql/examples/select-from-users.sql")
	if err != nil {
		return err
	}
	printAllUsersExample(result)

	result, err = database.SelectData(DataPath+"sql/examples/select-from-users-with-id.sql", 3)
	if err != nil {
		return err
	}
	printAllUsersExample(result)

	err = database.UpdateData(DataPath+"sql/examples/change-website-of-user.sql", "https://google.com", 3)
	if err != nil {
		return err
	}

	err = database.DeleteData(DataPath+"sql/examples/remove-from-users.sql", 5)
	if err != nil {
		return err
	}

	return nil
}

func run() error {
	// Initialize database
	database := new(Database)

	// Connect
	err := database.Connect()
	defer database.Connection.Close()
	if err != nil {
		return err
	}

	err = runExamples(database)
	if err != nil {
		return err
	}
	// We are done
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
