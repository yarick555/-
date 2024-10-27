package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string = "World" // создаем переменную, которую будем менять

type MessageRequest struct {
	Message string `json:"message"`
} //создаем json

func HelloHandler(w http.ResponseWriter, r *http.Request) { // get запрос
	fmt.Fprintf(w, "Hello, %s!", message) // результат гет запроса
}

func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} // пост запрос, изменение переменной
	message = req.Message
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")            // адрес гет
	router.HandleFunc("/api/message", UpdateMessageHandler).Methods("POST") // адрес пост

	http.ListenAndServe(":8080", router)
}
