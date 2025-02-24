package userService

import (
	"gorm.io/gorm"
	"project/internal/taskService"
)

type User struct {
	gorm.Model
	Email    string             `gorm:"unique;not null"` // Уникальный email
	Password string             `gorm:"not null"`        // Пароль
	Tasks    []taskService.Task `gorm:"foreignKey:UserID"`
}
