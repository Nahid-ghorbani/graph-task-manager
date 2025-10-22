package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model         // includes id, created_at, updated_at
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
}
