package controllers

import (
	"net/http"
	"task-manager/backend/database"
	"task-manager/backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTasks retrieves all tasks from the database
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// CreateTask adds a new task to the database
func CreateTask(c *gin.Context) {
	var task models.Task

	// Bind incoming JSON request to Task struct
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If DueDate is not set, set it to today
	if task.DueDate.IsZero() {
		task.DueDate = time.Now()
	}

	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// UpdateTask modifies an existing task
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask removes a task from the database
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
