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
func jsonResponse(w http.ResponseWriter, statusCode int, data any) {
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
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{Message: "Некорректный JSON"})
		return
	}

	// Простой валидатор: проверим, что имя и email не пустые
	if user.Name == "" || user.Email == "" {
		jsonResponse(w, http.StatusBadRequest, Response{Message: "Имя и Email обязательны"})
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	// Используем helper функцию вместо ручной установки заголовков
	jsonResponse(w, http.StatusCreated, user)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/welcome", home2Handler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
