package routes

import (
	"task-manager/backend/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up API routes
func RegisterRoutes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetTasks)          // Fetch all tasks
	router.POST("/tasks", controllers.CreateTask)       // Create a new task
	router.PUT("/tasks/:id", controllers.UpdateTask)    // Update a task
	router.DELETE("/tasks/:id", controllers.DeleteTask) // Delete a task
}
