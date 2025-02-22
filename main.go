package main

import (
	"fmt"
	"log"
	"task-manager/backend/database"
	"task-manager/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDatabase()

	// Create a Gin router
	router := gin.Default()

	// Serve static frontend files
	router.Static("/static", "./frontend")
	router.StaticFile("/", "./frontend/index.html") // Serve index.html as the homepage
	
	// Register API routes
	routes.RegisterRoutes(router)

	// Start the server
	port := ":8080"
	fmt.Println("🚀 Server running at http://localhost" + port)
	err := router.Run(port)
	if err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}
