package handlers

import (
	"context"
	"project/internal/taskService" // Импортируем наш сервис
	"project/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// NewHandler создает новый экземпляр Handler
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// PatchTasksId обновляет задачу по её ID
func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Получаем ID задачи из запроса
	taskID := request.Id

	// Получаем данные для обновления из тела запроса
	taskUpdate := request.Body

	// Создаем структуру для обновления задачи
	updatedTask := taskService.Task{
		ID:     taskID,
		Text:   *taskUpdate.Task,
		IsDone: *taskUpdate.IsDone,
	}

	// Обновляем задачу в сервисе
	updatedTask, err := h.Service.UpdateTaskByID(taskID, updatedTask)
	if err != nil {
		return nil, err
	}

	// Создаем структуру ответа
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем обновленную задачу
	return response, nil
}

// DeleteTasksId удаляет задачу по её ID
func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Получаем ID задачи из запроса
	taskID := request.Id

	// Удаляем задачу в сервисе
	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	// Возвращаем пустой ответ с кодом 204 (No Content)
	return tasks.DeleteTasksId204Response{}, nil
}
