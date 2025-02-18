package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
	"project/internal/userService"
	"project/internal/web/tasks"
	"project/internal/web/users"
)

func main() {
	database.InitDB()

	// Миграции для задач и пользователей
	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("Ошибка при автомиграции: %v", err)
	}

	// Репозитории и сервисы для задач
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	// Репозитории и сервисы для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userServices := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userServices)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация ручек для задач
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Регистрация ручек для пользователей
	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
