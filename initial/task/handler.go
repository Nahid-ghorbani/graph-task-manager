package task

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// Attach task routes to router
func (h Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks", h.GetAllTasks)
}

//Create new task
func (h Handler) CreateTask(c *gin.Context){
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		fmt.Println("****task:", &task)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

//Get all tasks
func (h Handler) GetAllTasks(c *gin.Context) {
	var tasks []Task
	
	if err := h.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No tasks found!!"})
	}

	c.JSON(http.StatusOK, tasks)
}

