package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Вторая задача с базой данных
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
	w.WriteHeader(http.StatusCreated)
}

func main() {
	InitDB()

	// Модель Task
	DB.AutoMigrate(&Task{})

	// Маршруты
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTaskHandler).Methods("POST")

	// Запускаем сервер
	http.ListenAndServe(":8080", router)
}
