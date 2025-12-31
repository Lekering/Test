package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура для пользователя
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Структура для ответа с ошибкой
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Хранилище пользователей (пока в памяти)
var users = []User{
	{ID: 1, Name: "Иван", Email: "ivan@example.com"},
	{ID: 2, Name: "Мария", Email: "maria@example.com"},
	{ID: 3, Name: "Петр", Email: "petr@example.com"},
}

// Получить всех пользователей
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")
	
	// Кодируем users в JSON и отправляем
	json.NewEncoder(w).Encode(users)
}

// Получить конкретного пользователя по ID
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Получаем ID из параметра: /user?id=1
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "bad_request",
			Message: "ID не указан",
		})
		return
	}
	
	// Ищем пользователя
	var foundUser *User
	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == idParam {
			foundUser = &user
			break
		}
	}
	
	if foundUser == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "not_found",
			Message: "Пользователь не найден",
		})
		return
	}
	
	json.NewEncoder(w).Encode(foundUser)
}

// Создать нового пользователя
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "method_not_allowed",
			Message: "Используйте POST",
		})
		return
	}
	
	// Декодируем JSON из тела запроса
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "invalid_json",
			Message: "Неверный формат JSON",
		})
		return
	}
	
	// Генерируем ID (в реальности это делает БД)
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/user/create", createUserHandler)
	
	fmt.Println("Сервер запущен на http://localhost:8080")
	fmt.Println("\nПопробуйте:")
	fmt.Println("GET  http://localhost:8080/users")
	fmt.Println("GET  http://localhost:8080/user?id=1")
	fmt.Println("POST http://localhost:8080/user/create")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
ЗАДАНИЯ НА ДЕНЬ 16:
1. Запустите сервер

2. Проверьте GET запросы в браузере:
   - http://localhost:8080/users
   - http://localhost:8080/user?id=2

3. Для POST используйте curl или Postman:
   curl -X POST http://localhost:8080/user/create \
     -H "Content-Type: application/json" \
     -d '{"name":"Анна","email":"anna@example.com"}'

4. ПРАКТИКА:
   - Добавьте поле Age в User
   - Создайте endpoint для удаления пользователя
   - Добавьте валидацию (email должен содержать @)
*/
