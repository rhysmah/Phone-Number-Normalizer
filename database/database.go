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
	err := db.Close()
	if err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}

	return nil
}

// CreateTable creates the "phoneNumbers" table in the SQLite database
// if it does not already exist.
//
// The table schema includes:
// - id: an INTEGER PRIMARY KEY with AUTOINCREMENT
// - numbers: a TEXT field
//
// Returns:
// - nil if the table is successfully created or already exists.
// - An error if there is an issue executing the table creation.
func CreateTable(numbers []string) error {
	createTable := `CREATE TABLE IF NOT EXISTS phoneNumbers (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		numbers TEXT
	);`

	_, err := db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	if err := insertInitialData(numbers); err != nil {
		return fmt.Errorf("error inserting initial data into table: %w", err)
	}

	return nil
}

func insertInitialData(numbers []string) error {
	var count int

	err := db.QueryRow(`SELECT COUNT(*) from phoneNumbers`).Scan(&count)
	if err != nil {
		return fmt.Errorf("error accessing data in database: %w", err)
	}

	if count != 0 {
		log.Println("Database already contains initial data")
		return nil
	}

	for _, number := range numbers {
		_, err := db.Exec(`INSERT INTO phoneNumbers (numbers) VALUES (?)`, number)
		if err != nil {
			return fmt.Errorf("error inserting number '%s' column in table: %w", number, err)
		}
	}
	log.Println("Initial data successfully added to database")
	return nil
}

func ReadNumbers() ([]string, error) {
	var numbers []string

	q := "SELECT * FROM phoneNumbers"

	// Extract rows that meet query
	rows, err := db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("error reading data: %w", err)
	}
	defer rows.Close()

	// Scan each row, read numbers, add them to string slice
	for rows.Next() {

		var id int64
		var num string

		if err := rows.Scan(&id, &num); err != nil {
			return nil, fmt.Errorf("error reading rows: %w", err)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
