package userService

import "gorm.io/gorm"

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUsers() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) UpdateUser(id uint, user *User) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}
