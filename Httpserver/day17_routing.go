package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var users = map[int]User{
	1: {ID: 1, Name: "Иван", Email: "ivan@example.com"},
	2: {ID: 2, Name: "Мария", Email: "maria@example.com"},
}
var nextID = 3

// Общий обработчик для /users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r)
	case http.MethodPost:
		handleCreateUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Метод не поддерживается"})
	}
}

// Обработчик для /users/{id}
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Извлекаем ID из пути: /users/123
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Неверный ID"})
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r, id)
	case http.MethodPut:
		handleUpdateUser(w, r, id)
	case http.MethodDelete:
		handleDeleteUser(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Метод не поддерживается"})
	}
}

// GET /users - получить всех пользователей
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	usersList := make([]User, 0, len(users))
	for _, user := range users {
		usersList = append(usersList, user)
	}
	json.NewEncoder(w).Encode(usersList)
}

// POST /users - создать пользователя
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Неверный JSON"})
		return
	}
	
	newUser.ID = nextID
	nextID++
	users[newUser.ID] = newUser
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// GET /users/{id} - получить пользователя
func handleGetUser(w http.ResponseWriter, r *http.Request, id int) {
	user, exists := users[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Пользователь не найден"})
		return
	}
	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id} - обновить пользователя
func handleUpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	if _, exists := users[id]; !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Пользователь не найден"})
		return
	}
	
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Неверный JSON"})
		return
	}
	
	updatedUser.ID = id
	users[id] = updatedUser
	
	json.NewEncoder(w).Encode(updatedUser)
}

// DELETE /users/{id} - удалить пользователя
func handleDeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	if _, exists := users[id]; !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Пользователь не найден"})
		return
	}
	
	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}

// Простой роутер
func router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		usersHandler(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/users/") {
		userHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Путь не найден"})
	}
}

func main() {
	http.HandleFunc("/", router)
	
	fmt.Println("🚀 REST API запущен на http://localhost:8080")
	fmt.Println("\nДоступные endpoints:")
	fmt.Println("GET    /users          - Все пользователи")
	fmt.Println("POST   /users          - Создать пользователя")
	fmt.Println("GET    /users/{id}     - Получить пользователя")
	fmt.Println("PUT    /users/{id}     - Обновить пользователя")
	fmt.Println("DELETE /users/{id}     - Удалить пользователя")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
ЗАДАНИЯ НА ДЕНЬ 17:
1. Тестируйте с помощью curl:
   
   # Получить всех
   curl http://localhost:8080/users
   
   # Создать
   curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name":"Анна","email":"anna@example.com"}'
   
   # Получить одного
   curl http://localhost:8080/users/1
   
   # Обновить
   curl -X PUT http://localhost:8080/users/1 \
     -H "Content-Type: application/json" \
     -d '{"name":"Иван Иванов","email":"ivan.ivanov@example.com"}'
   
   # Удалить
   curl -X DELETE http://localhost:8080/users/1

2. ПРАКТИКА:
   - Добавьте endpoint для поиска: GET /users/search?name=Иван
   - Добавьте PATCH для частичного обновления
   - Создайте ресурс "Posts" с аналогичным CRUD
*/
