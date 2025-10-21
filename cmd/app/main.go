package main

import (
	"log"

	"github.com/Nahid-ghorbani/graph-task-manager/initial/db"
)

func main() {

	database := db.Connect()
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate! , error : %v", err)
	}
	log.Println("Migration completed!")
}
