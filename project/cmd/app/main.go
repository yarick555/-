package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/update/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
