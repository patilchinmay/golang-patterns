package main

import (
	"log"
	"time"
)

type database struct {
	name      string
	createdAt time.Time
}

// db is a pointer to database. So that we only have one instance.
var db *database

// databaseCreator works as a singleton implementor
func databaseCreator() *database {
	if db != nil {
		return db
	} else {
		db = &database{
			name:      "Database Uno",
			createdAt: time.Now(),
		}
		return db
	}
}

func main() {
	databaseInstance := databaseCreator()
	log.Printf("databaseInstance : %v\n\n", databaseInstance)

	databaseInstance2 := databaseCreator()
	log.Printf("databaseInstance2 : %v\n\n", databaseInstance2)

	log.Printf("databaseInstance == databaseInstance2 : %v", databaseInstance == databaseInstance2)
}
