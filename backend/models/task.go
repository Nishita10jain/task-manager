package models

import "time"

// Task represents a task in the database
type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	DueDate     time.Time `json:"due_date"`
}
