package main

import (
	"fmt"
	"num-normalizer/database"
	"num-normalizer/file"
)

func main() {
	var err error

	phoneNumbers, err := file.ProcessFile()
	if err != nil {
		fmt.Println(err)
	}

	if err = database.CreateDatabase(phoneNumbers); err != nil {
		fmt.Println(err)
	}

	err = database.NormalizeAndUpdateNumbersInDB()
	if err != nil {
		fmt.Println(err)
	}
}
