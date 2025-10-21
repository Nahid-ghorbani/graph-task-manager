package main

import (
	"log"

	"github.com/Nahid-ghorbani/graph-task-manager/initial/db"
	"github.com/Nahid-ghorbani/graph-task-manager/initial/task"
)

func main() {

	database := db.Connect()
	if err := database.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("Failed to migrate! , error : %v", err)
	}
	log.Println("Migration completed!")
}
