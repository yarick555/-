package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"` // Уникальный email
	Password string `gorm:"not null"`        // Пароль
}
