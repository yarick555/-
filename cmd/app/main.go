package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
	"project/internal/web/tasks"
)

func main() {
	// Инициализация
	database.InitDB()

	// Автомиграция
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("Ошибка при автомиграции: %v", err)
	}

	// Создаем репу сервис и хендлер
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем хендлеры
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
