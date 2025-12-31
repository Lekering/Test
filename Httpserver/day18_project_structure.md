# День 18-19: Правильная структура проекта

## Структура папок

```
myapi/
├── main.go              # Точка входа
├── go.mod               # Зависимости
├── internal/            # Приватный код (не импортируется извне)
│   ├── handlers/        # HTTP обработчики
│   │   └── user.go
│   ├── models/          # Структуры данных
│   │   └── user.go
│   ├── storage/         # Работа с хранилищем
│   │   └── memory.go
│   └── middleware/      # Middleware функции
│       └── logger.go
└── pkg/                 # Публичный код (может импортироваться)
    └── response/        # Утилиты для ответов
        └── response.go
```

## main.go

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	
	"myapi/internal/handlers"
	"myapi/internal/middleware"
	"myapi/internal/storage"
)

func main() {
	// Инициализируем хранилище
	store := storage.NewMemoryStorage()
	
	// Создаем обработчики
	userHandler := handlers.NewUserHandler(store)
	
	// Создаем роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.HandleUsers)
	mux.HandleFunc("/users/", userHandler.HandleUser)
	
	// Добавляем middleware
	handler := middleware.Logger(mux)
	
	fmt.Println("🚀 Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
```

## internal/models/user.go

```go
package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !strings.Contains(r.Email, "@") {
		return fmt.Errorf("invalid email")
	}
	return nil
}
```

## internal/storage/memory.go

```go
package storage

import (
	"errors"
	"sync"
	
	"myapi/internal/models"
)

var ErrNotFound = errors.New("not found")

type Storage interface {
	GetAll() []models.User
	GetByID(id int) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(id int, user models.User) (*models.User, error)
	Delete(id int) error
}

type MemoryStorage struct {
	mu     sync.RWMutex
	users  map[int]models.User
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (s *MemoryStorage) GetAll() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	result := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		result = append(result, user)
	}
	return result
}

func (s *MemoryStorage) GetByID(id int) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	user, exists := s.users[id]
	if !exists {
		return nil, ErrNotFound
	}
	return &user, nil
}

func (s *MemoryStorage) Create(user models.User) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return &user, nil
}

func (s *MemoryStorage) Update(id int, user models.User) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.users[id]; !exists {
		return nil, ErrNotFound
	}
	
	user.ID = id
	s.users[id] = user
	return &user, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.users[id]; !exists {
		return ErrNotFound
	}
	
	delete(s.users, id)
	return nil
}
```

## internal/handlers/user.go

```go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	
	"myapi/internal/models"
	"myapi/internal/storage"
	"myapi/pkg/response"
)

type UserHandler struct {
	store storage.Storage
}

func NewUserHandler(store storage.Storage) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *UserHandler) HandleUser(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		h.getUser(w, r, id)
	case http.MethodPut:
		h.updateUser(w, r, id)
	case http.MethodDelete:
		h.deleteUser(w, r, id)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users := h.store.GetAll()
	response.JSON(w, http.StatusOK, users)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid json")
		return
	}
	
	if err := req.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	
	created, err := h.store.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create user")
		return
	}
	
	response.JSON(w, http.StatusCreated, created)
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request, id int) {
	user, err := h.store.GetByID(id)
	if err == storage.ErrNotFound {
		response.Error(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get user")
		return
	}
	
	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid json")
		return
	}
	
	updated, err := h.store.Update(id, user)
	if err == storage.ErrNotFound {
		response.Error(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to update user")
		return
	}
	
	response.JSON(w, http.StatusOK, updated)
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	err := h.store.Delete(id)
	if err == storage.ErrNotFound {
		response.Error(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to delete user")
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

func extractID(path string) (int, error) {
	parts := strings.Split(strings.TrimPrefix(path, "/users/"), "/")
	if len(parts) == 0 {
		return 0, errors.New("invalid path")
	}
	return strconv.Atoi(parts[0])
}
```

## pkg/response/response.go

```go
package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, ErrorResponse{Error: message})
}
```

## internal/middleware/logger.go

```go
package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Вызываем следующий обработчик
		next.ServeHTTP(w, r)
		
		// Логируем после обработки
		log.Printf(
			"%s %s %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
```

## Инициализация проекта

```bash
# Создайте структуру папок
mkdir -p myapi/internal/{handlers,models,storage,middleware}
mkdir -p myapi/pkg/response

# Инициализируйте Go модуль
cd myapi
go mod init myapi

# Создайте все файлы из примеров выше
# Затем запустите
go run main.go
```

## Преимущества такой структуры

1. **Разделение ответственности** - каждый пакет отвечает за свое
2. **Тестируемость** - легко писать тесты для каждого слоя
3. **Расширяемость** - легко добавить новые ресурсы
4. **Читаемость** - понятно, где что находится
5. **Безопасность** - internal/ нельзя импортировать извне

## Задания

1. Реализуйте все файлы из примера
2. Добавьте ресурс "Posts" по аналогии с User
3. Добавьте middleware для CORS
4. Добавьте валидацию для всех полей
