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
