package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Третья задача CRUD
// Переменная для хранения
var DB *gorm.DB

// Инииализация
func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

// Структурка для хранения
type Task struct {
	gorm.Model
	Task   string `json:"task"`    // Наш сервер будет ожидать json c полем text
	IsDone bool   `json:"is_done"` // В GO используем CamelCase, в Json - snake
}

// GET
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	// Получаем записи из базы
	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Возвращаем список задач в джсон
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// POST
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Декодируем пришедший джсон
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Новая запись в базе
	if err := DB.Create(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Создали успешно
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// PATCH
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Декодируем
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем поля
	if updatedTask.Task != "" {
		task.Task = updatedTask.Task
	}
	task.IsDone = updatedTask.IsDone

	if err := DB.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DELETE
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := DB.Delete(&Task{}, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	InitDB()

	// Модель Task
	DB.AutoMigrate(&Task{})

	// Маршруты
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteTaskHandler).Methods("DELETE")

	// Запускаем сервер
	http.ListenAndServe(":8080", router)
}
