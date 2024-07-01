package main

import (
	"fmt"
	"num-normalizer/database"
	"num-normalizer/file"
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
}
