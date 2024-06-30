package main

import (
	"num-normalizer/database"
)

func main() {
	err := database.InitializeDB()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB()

}
