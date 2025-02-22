package database

import (
	"log"
	"task-manager/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global DB variable
var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto-migrate database schema
	DB.AutoMigrate(&models.Task{})

	log.Println("✅ Database initialized successfully!")
}
