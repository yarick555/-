package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
	GetTasksForUser(userID uint) ([]Task, error) // Добавляем новый метод
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// CreateTask создает новую задачу в базе данных
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

// GetAllTasks возвращает все задачи из базы данных
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// UpdateTaskByID обновляет задачу по её ID
func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	// Находим задачу по ID и обновляем её поля
	result := r.db.Model(&Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	// Возвращаем обновленную задачу
	var updatedTask Task
	err := r.db.First(&updatedTask, id).Error
	if err != nil {
		return Task{}, err
	}
	return updatedTask, nil
}

// DeleteTaskByID удаляет задачу по её ID
func (r *taskRepository) DeleteTaskByID(id uint) error {
	// Удаляем задачу по ID
	result := r.db.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *taskRepository) GetTasksForUser(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
