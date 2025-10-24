package main

import (
	"log"

	"github.com/Nahid-ghorbani/graph-task-manager/internal/db"
	"github.com/Nahid-ghorbani/graph-task-manager/internal/task"
	"github.com/gin-gonic/gin"
)

func main() {

	database := db.Connect()
	if err := database.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("Failed to migrate! , error : %v", err)
	}
	log.Println("Migration completed!")

	router := gin.Default()

	taskRepo := task.NewGormRepository(database)
	taskHandler := task.NewTaskHandler(taskRepo)
	taskHandler.RegisterRoutes(router)

	router.Run(":8080")
}
