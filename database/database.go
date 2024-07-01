package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitializeDB opens a connection to the SQLite database "database.db"
// and verifies the connection by pinging the database.
//
// Returns:
// - nil if the connection is successfully established and verified.
// - An error if there is an issue opening or pinging the database.
func InitializeDB() error {
	var err error

	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("error initializing database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}

// CloseDB closes the connection to the SQLite database.
//
// Returns:
// - nil if the connection is successfully closed.
// - An error if there is an issue closing the database connection.
func CloseDB() error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}

	return nil
}

// CreateTable creates the "phoneNumbers" table in the SQLite database
// if it does not already exist and inserts initial data if the table is empty.
//
// The table schema includes:
// - id: an INTEGER PRIMARY KEY with AUTOINCREMENT
// - numbers: a TEXT field
//
// Parameters:
// - numbers: A slice of strings containing the initial data to be inserted.
//
// Returns:
// - nil if the table is successfully created or already exists and data is inserted if the table is empty.
// - An error if there is an issue executing the table creation or inserting the initial data.
func CreateTable(numbers []string) error {
	createTable := `CREATE TABLE IF NOT EXISTS phoneNumbers (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		numbers TEXT
	);`

	if _, err := db.Exec(createTable); err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	if err := insertInitialData(numbers); err != nil {
		return fmt.Errorf("error inserting initial data into table: %w", err)
	}

	return nil
}

// insertInitialData inserts the initial data into the "phoneNumbers" table
// if the table is empty.
//
// Parameters:
// - numbers: A slice of strings containing the initial data to be inserted.
//
// Returns:
// - nil if the initial data is successfully inserted or if the table is already populated.
// - An error if there is an issue accessing the data or inserting the initial data.
func insertInitialData(numbers []string) error {
	var count int

	if err := db.QueryRow(`SELECT COUNT(*) from phoneNumbers`).Scan(&count); err != nil {
		return fmt.Errorf("error accessing data in database: %w", err)
	}

	if count != 0 {
		log.Println("Database already contains initial data")
		return nil
	}

	for _, number := range numbers {
		if _, err := db.Exec(`INSERT INTO phoneNumbers (numbers) VALUES (?)`, number); err != nil {
			return fmt.Errorf("error inserting number '%s' column in table: %w", number, err)
		}
	}
	log.Println("Initial data successfully added to database")
	return nil
}

// ReadNumbers reads all phone numbers from the "phoneNumbers" table.
//
// Returns:
// - A slice of strings containing the phone numbers.
// - An error if there is an issue reading the data.
func ReadNumbers() ([]string, error) {
	var numbers []string

	rows, err := db.Query("SELECT * FROM phoneNumbers")
	if err != nil {
		return nil, fmt.Errorf("error reading data: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var num string
		if err := rows.Scan(&num); err != nil {
			return nil, fmt.Errorf("error reading rows: %w", err)
		}
		numbers = append(numbers, num)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return numbers, nil
}
