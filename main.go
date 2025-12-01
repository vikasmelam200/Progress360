package main

import (
	"fmt"

	"github.com/vikasmelam200/Progress360/db"
)

func main() {
	// starts with database connection:
	dbUri := "mongodb://localhost:27017"
	dbName := "Progress360"
	client, db, err := db.ConnectDB(dbUri, dbName)
	fmt.Println(client)
	fmt.Println(db.Name())
	if err != nil {
		fmt.Println(err)
	}
}
