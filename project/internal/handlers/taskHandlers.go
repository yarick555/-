package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/taskService" // Импортируем наш сервис
	"strconv"
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

// GetTasksHandler возвращает все задачи
func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// PostTaskHandler создает новую задачу
func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

// UpdateTaskHandler обновляет задачу по ID
func (h *Handler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи из URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Декодируем тело запроса в структуру Task
	var task taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем задачу
	updatedTask, err := h.Service.UpdateTaskByID(uint(id), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем обновленную задачу
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTaskHandler удаляет задачу по ID
func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи из URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Удаляем задачу
	if err := h.Service.DeleteTaskByID(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем статус 204 (No Content)
	w.WriteHeader(http.StatusNoContent)
}
