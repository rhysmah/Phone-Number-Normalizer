package main

import (
	"fmt"
	"num-normalizer/database"
	"num-normalizer/file"
	"num-normalizer/processor"
)

func main() {
	var err error

	if err = database.InitializeDB(); err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := database.CloseDB(); err != nil {
			fmt.Println(err)
		}
	}()

	phoneNumbers, err := file.ProcessFile()
	if err != nil {
		fmt.Println(err)
	}

	if err = database.CreateTable(phoneNumbers); err != nil {
		fmt.Println(err)
	}

	normalizedData, err := processor.NormalizeNumbers(phoneNumbers)
	if err != nil {
		fmt.Println(err)
	}

	if err := database.InsertData(normalizedData); err != nil {
		fmt.Println(err)
	}
}
