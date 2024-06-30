package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB() error {
	var err error

	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("error initializing database: %w", err)
	}

	// Ensure database is accessible
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}

func CloseDB() error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}

	return nil
}

func CreateTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS phoneNumbers (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		numbers TEXT
	);`

	_, err := db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	return nil
}

func AddNumbers(numbers []string) error {
	for _, number := range numbers {
		_, err := db.Exec(`INSERT INTO phoneNumbers (numbers) VALUES (?)`, number)
		if err != nil {
			return fmt.Errorf("error inserting number '%s' column in table: %w", number, err)
		}
	}
	return nil
}
