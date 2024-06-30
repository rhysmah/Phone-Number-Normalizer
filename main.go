package main

import (
	"fmt"
	"num-normalizer/database"
	"num-normalizer/file"
)

func main() {
	var err error

	// Read file contents and store them in string slice
	phoneNumbers, err := file.ProcessFile()
	if err != nil {
		fmt.Println(err)
	}

	// Initialize database; if already exists, access it
	if err = database.InitializeDB(); err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := database.CloseDB(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create table to store read phoneNumbers ([]string)
	if err = database.CreateTable(); err != nil {
		fmt.Println(err)
	}

	// Store phone numbers in database
	if err = database.AddNumbers(phoneNumbers); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully added numbers to database")
}
