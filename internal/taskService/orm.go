package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	ID     uint   `json:"id"`
	Text   string `json:"text"`
}
