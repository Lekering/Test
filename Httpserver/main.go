package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func home2Handler(w http.ResponseWriter, r *http.Request) {
	// Эта строка устанавливает заголовок ответа Content-Type в значение "application/json",
	// чтобы клиент понял, что ответ содержит JSON-данные
	w.Header().Set("Content-Type", "application/json")
	// Эта строка кодирует структуру Response в JSON и отправляет клиенту как ответ.
	json.NewEncoder(w).Encode(Response{Message: "Добро пожаловать в API"})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет! Это мой первый HTTP сервер на Go!")

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home2Handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
