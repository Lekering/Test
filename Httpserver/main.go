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

// jsonResponse устанавливает Content-Type и кодирует данные в JSON
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func home2Handler(w http.ResponseWriter, r *http.Request) {
	// Используем helper функцию для отправки JSON ответа
	jsonResponse(w, http.StatusOK, Response{Message: "Добро пожаловать в API"})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет! Это мой первый HTTP сервер на Go!")

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	// json.NewDecoder(r.Body).Decode(&user) - создает новый декодер, который читает JSON из тела запроса r.Body,
	// и декодирует его в структуру user (заполняет ее полями из запроса).
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)

	// Используем helper функцию вместо ручной установки заголовков
	jsonResponse(w, http.StatusCreated, user)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", createUserHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
