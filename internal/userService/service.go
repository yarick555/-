package userService

import "project/internal/taskService"

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(email, password string) (*User, error) {
	user := &User{Email: email, Password: password}
	err := s.Repo.CreateUser(user)
	return user, err
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.Repo.GetUsers()
}

func (s *UserService) UpdateUser(id uint, email, password string) error {
	user := &User{Email: email, Password: password}
	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := s.Repo.DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
