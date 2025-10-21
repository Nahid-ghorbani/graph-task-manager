package main

import (
	"log"

	"github.com/Nahid-ghorbani/graph-task-manager/initial/db"
	"github.com/Nahid-ghorbani/graph-task-manager/initial/task"
	"github.com/gin-gonic/gin"
)

func main() {

	database := db.Connect()
	if err := database.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("Failed to migrate! , error : %v", err)
	}
	log.Println("Migration completed!")

	router := gin.Default()

	taskHandler := task.Handler{DB: database}
	taskHandler.RegisterRoutes(router)

	router.Run(":8080")
}
