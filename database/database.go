package database

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize database and table with initial data

func CreateDatabase(initialData []string) error {
	if err := initializeDB(); err != nil {
		return err
	}

	defer func() {
		if err := closeDB(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}()

	if err := createTable(); err != nil {
		return err
	}

	if err := insertInitialData(initialData); err != nil {
		return err
	}

	log.Printf("successfully initialized database with initial data")
	return nil
}

func initializeDB() error {
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

func closeDB() error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}

	return nil
}

func createTable() error {

	createTable := `CREATE TABLE IF NOT EXISTS phoneNumbers (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		numbers TEXT
	);`

	if _, err := db.Exec(createTable); err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	return nil
}

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

	return nil
}

// Normalize data and remove duplicates

func NormalizeAndUpdateNumbersInDB() error {

	err := initializeDB()
	if err != nil {
		return err
	}

	defer func() error {
		if err := closeDB(); err != nil {
			return err
		}
		return nil
	}()

	rawNumbers, err := readNumbers()
	if err != nil {
		return fmt.Errorf("error reading numbers in database: %w", err)
	}

	normalizedNumbers, err := processNumbers(rawNumbers)
	if err != nil {
		return fmt.Errorf("error normalizing numbers: %w", err)
	}

	if err := updateNumbersInDB(normalizedNumbers); err != nil {
		return fmt.Errorf("error updating databse: %w", err)
	}

	log.Println("Successfully updated database with normalized numbers")
	return nil
}

func readNumbers() (map[int]string, error) {
	phoneNumbers := make(map[int]string)

	rows, err := db.Query("SELECT id, numbers FROM phoneNumbers")
	if err != nil {
		return nil, fmt.Errorf("error reading data in table: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var num string

		if err := rows.Scan(&id, &num); err != nil {
			return nil, fmt.Errorf("error reading rows: %w", err)
		}
		phoneNumbers[id] = num
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	log.Println("Successfully read all numbers from database")
	return phoneNumbers, nil
}

func processNumbers(rawNumbers map[int]string) (map[int]string, error) {
	normalizedNumbers := make(map[int]string)

	for id, number := range rawNumbers {
		normalized, err := normalizeNumber(number)
		if err != nil {
			log.Printf("error normalizing number %s with id %d: %v", number, id, err)
			continue
		}

		normalizedNumbers[id] = normalized // removes duplicates
	}

	log.Printf("Successfully normalized all phone numbers and removed duplicates")
	return normalizedNumbers, nil
}

func normalizeNumber(rawNumber string) (string, error) {
	regex, err := regexp.Compile(`\D`)
	if err != nil {
		return "", fmt.Errorf("error parsing regular expression: %w", err)
	}

	normalized := regex.ReplaceAllString(rawNumber, "")
	return normalized, nil
}

func updateNumbersInDB(normalizedNumbers map[int]string) error {
	alreadyExists := make(map[string]bool)

	for id, number := range normalizedNumbers {
		if alreadyExists[number] {
			if _, err := db.Exec("DELETE FROM phoneNumbers WHERE id = ?", id); err != nil {
				return fmt.Errorf("error deleting duplicate phone number %s with id %d from database", number, id)
			}
		} else {
			if _, err := db.Exec("UPDATE phoneNumbers SET numbers = ? WHERE id = ?", number, id); err != nil {
				return fmt.Errorf("error updating normalized phone number for id '%d': %w", id, err)
			}
		}
		alreadyExists[number] = true
	}

	log.Println("Successfully updated numbers in database")
	return nil
}
