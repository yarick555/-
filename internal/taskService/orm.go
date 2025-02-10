package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`    // Краткое название задачи
	IsDone bool   `json:"is_done"` // Статус выполнения (по умолчанию false)
}
