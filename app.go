package main

import (
	"context"
	"indexer/storage"
	"log"
	"os"
)

// Database file name
const DATABASE = "data/data.db"

var Database *storage.Database

func main() {

	// Init the database
	var err error
	Database, err = storage.NewDatabase(DATABASE)
	if err != nil {
		panic(err)
	}
	defer Database.Close() // Close it at the end

	cmd := InitCli()
	// Run the cmd
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
