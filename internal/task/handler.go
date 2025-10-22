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
	r.DELETE("/tasks/:id", h.DeleteTask)
	r.GET("/tasks/:id", h.GetTaskDetail)
	r.PATCH("/tasks/:id", h.UpdateTask)
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

//Delete task using id
func (h Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task Task

	if err := h.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No task found with this id"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}


//Get task detail using id
func (h Handler) GetTaskDetail(c *gin.Context) {
	id := c.Param("id")

	var task Task
	if err := h.DB.First(&task , id).Error ; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No task found with this id"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// update task using patch using id
func (h Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := h.DB.First(&task, id).Error ; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No task found with this id"})
			return 
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	} 

	if err := h.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, task)
}