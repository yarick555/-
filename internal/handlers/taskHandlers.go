package handlers

import (
	"context"
	"errors"
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
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,     // Добавляем ID
			Task:   &tsk.Task,   // Добавляем Task
			IsDone: &tsk.IsDone, // Добавляем IsDone
		}
		response = append(response, task)
	}

	// Возвращаем ответ
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Проверяем, что поле Task не nil
	if request.Body.Task == nil {
		return nil, errors.New("task field is required")
	}

	// Создаем задачу с полем Task и IsDone по умолчанию (false)
	taskToCreate := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: false, // По умолчанию задача не выполнена
	}

	// Обращаемся к сервису и создаем задачу
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// Создаем структуру ответа
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,     // Возвращаем ID
		Task:   &createdTask.Task,   // Возвращаем Task
		IsDone: &createdTask.IsDone, // Возвращаем IsDone
	}

	// Возвращаем ответ
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
		Task:   *taskUpdate.Task,
		IsDone: *taskUpdate.IsDone,
	}

	// Обновляем задачу в сервисе
	updatedTask, err := h.Service.UpdateTaskByID(taskID, updatedTask)
	if err != nil {
		return nil, err
	}

	// Создаем структуру ответа
	response := tasks.PatchTasksId200JSONResponse{
		Task:   &updatedTask.Task,
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
