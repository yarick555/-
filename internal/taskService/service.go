package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask создает новую задачу
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

// GetAllTasks возвращает все задачи из базы данных
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

// UpdateTaskByID обновляет задачу по её ID
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

// DeleteTaskByID удаляет задачу по её ID
func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

// GetTasksForUser возвращает все задачи для конкретного пользователя
func (s *TaskService) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksForUser(userID)
}
