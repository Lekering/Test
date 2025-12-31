# День 22-28: Финальный проект - Blog API

Создадим полноценный REST API для блога с постами, комментариями и пользователями.

## Структура проекта

```
blog-api/
├── main.go
├── go.mod
├── internal/
│   ├── models/
│   │   ├── user.go
│   │   ├── post.go
│   │   └── comment.go
│   ├── storage/
│   │   └── memory.go
│   ├── handlers/
│   │   ├── user.go
│   │   ├── post.go
│   │   └── comment.go
│   ├── middleware/
│   │   ├── logger.go
│   │   ├── auth.go
│   │   └── cors.go
│   └── router/
│       └── router.go
└── pkg/
    └── response/
        └── response.go
```

## Шаг 1: Инициализация (День 22)

```bash
mkdir blog-api
cd blog-api
go mod init blog-api

mkdir -p internal/{models,storage,handlers,middleware,router}
mkdir -p pkg/response
```

## Шаг 2: Модели (internal/models/)

### user.go
```go
package models

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Username == "" {
		return errors.New("username is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" || len(r.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}
```

### post.go
```go
package models

import (
	"errors"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *CreatePostRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.Content == "" {
		return errors.New("content is required")
	}
	if len(r.Title) > 200 {
		return errors.New("title too long")
	}
	return nil
}
```

### comment.go
```go
package models

import (
	"errors"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
}

func (r *CreateCommentRequest) Validate() error {
	if r.Content == "" {
		return errors.New("content is required")
	}
	if len(r.Content) > 1000 {
		return errors.New("comment too long")
	}
	return nil
}
```

## Шаг 3: Storage (internal/storage/memory.go)

```go
package storage

import (
	"errors"
	"sync"
	"time"

	"blog-api/internal/models"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrUnauthorized  = errors.New("unauthorized")
)

type Storage struct {
	mu           sync.RWMutex
	users        map[int]*models.User
	posts        map[int]*models.Post
	comments     map[int]*models.Comment
	userNextID   int
	postNextID   int
	commentNextID int
}

func NewStorage() *Storage {
	return &Storage{
		users:    make(map[int]*models.User),
		posts:    make(map[int]*models.Post),
		comments: make(map[int]*models.Comment),
		userNextID:   1,
		postNextID:   1,
		commentNextID: 1,
	}
}

// ===== USER METHODS =====

func (s *Storage) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &models.User{
		ID:        s.userNextID,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	s.userNextID++
	s.users[user.ID] = user
	return user, nil
}

func (s *Storage) GetUser(id int) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, ErrNotFound
	}
	return user, nil
}

func (s *Storage) GetAllUsers() []*models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// ===== POST METHODS =====

func (s *Storage) CreatePost(userID int, req models.CreatePostRequest) (*models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Проверяем, что пользователь существует
	if _, exists := s.users[userID]; !exists {
		return nil, ErrNotFound
	}

	post := &models.Post{
		ID:        s.postNextID,
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.postNextID++
	s.posts[post.ID] = post
	return post, nil
}

func (s *Storage) GetPost(id int) (*models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, exists := s.posts[id]
	if !exists {
		return nil, ErrNotFound
	}
	return post, nil
}

func (s *Storage) GetAllPosts() []*models.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts := make([]*models.Post, 0, len(s.posts))
	for _, post := range s.posts {
		posts = append(posts, post)
	}
	return posts
}

func (s *Storage) UpdatePost(postID, userID int, req models.CreatePostRequest) (*models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[postID]
	if !exists {
		return nil, ErrNotFound
	}

	// Проверяем, что пост принадлежит пользователю
	if post.UserID != userID {
		return nil, ErrUnauthorized
	}

	post.Title = req.Title
	post.Content = req.Content
	post.UpdatedAt = time.Now()
	return post, nil
}

func (s *Storage) DeletePost(postID, userID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[postID]
	if !exists {
		return ErrNotFound
	}

	if post.UserID != userID {
		return ErrUnauthorized
	}

	delete(s.posts, postID)
	return nil
}

// ===== COMMENT METHODS =====

func (s *Storage) CreateComment(postID, userID int, req models.CreateCommentRequest) (*models.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Проверяем, что пост существует
	if _, exists := s.posts[postID]; !exists {
		return nil, ErrNotFound
	}

	comment := &models.Comment{
		ID:        s.commentNextID,
		PostID:    postID,
		UserID:    userID,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	s.commentNextID++
	s.comments[comment.ID] = comment
	return comment, nil
}

func (s *Storage) GetPostComments(postID int) []*models.Comment {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comments := make([]*models.Comment, 0)
	for _, comment := range s.comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments
}
```

## Шаг 4: Response helper (pkg/response/response.go)

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
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, ErrorResponse{Error: message})
}
```

## Шаг 5: Handlers

Создайте обработчики для пользователей, постов и комментариев по аналогии с предыдущими примерами.

## Шаг 6: Router (internal/router/router.go)

```go
package router

import (
	"net/http"

	"blog-api/internal/handlers"
	"blog-api/internal/middleware"
	"blog-api/internal/storage"
)

func New(store *storage.Storage) http.Handler {
	mux := http.NewServeMux()

	// Создаем handlers
	userHandler := handlers.NewUserHandler(store)
	postHandler := handlers.NewPostHandler(store)
	commentHandler := handlers.NewCommentHandler(store)

	// Регистрируем routes
	mux.HandleFunc("/users", userHandler.HandleUsers)
	mux.HandleFunc("/users/", userHandler.HandleUser)

	mux.HandleFunc("/posts", postHandler.HandlePosts)
	mux.HandleFunc("/posts/", postHandler.HandlePost)

	mux.HandleFunc("/posts/{id}/comments", commentHandler.HandleComments)

	// Применяем middleware
	var handler http.Handler = mux
	handler = middleware.CORS(handler)
	handler = middleware.Logger(handler)
	handler = middleware.Recovery(handler)

	return handler
}
```

## Шаг 7: main.go

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"blog-api/internal/router"
	"blog-api/internal/storage"
)

func main() {
	// Инициализируем storage
	store := storage.NewStorage()

	// Создаем router
	handler := router.New(store)

	// Запускаем сервер
	addr := ":8080"
	fmt.Printf("🚀 Blog API запущен на http://localhost%s\n", addr)
	fmt.Println("\nEndpoints:")
	fmt.Println("  POST   /users           - Создать пользователя")
	fmt.Println("  GET    /users           - Все пользователи")
	fmt.Println("  GET    /users/{id}      - Получить пользователя")
	fmt.Println("")
	fmt.Println("  POST   /posts           - Создать пост")
	fmt.Println("  GET    /posts           - Все посты")
	fmt.Println("  GET    /posts/{id}      - Получить пост")
	fmt.Println("  PUT    /posts/{id}      - Обновить пост")
	fmt.Println("  DELETE /posts/{id}      - Удалить пост")
	fmt.Println("")
	fmt.Println("  POST   /posts/{id}/comments - Добавить комментарий")
	fmt.Println("  GET    /posts/{id}/comments - Комментарии к посту")

	log.Fatal(http.ListenAndServe(addr, handler))
}
```

## Тестирование API

```bash
# Создать пользователя
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username":"john","email":"john@example.com","password":"secret123"}'

# Создать пост
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -d '{"title":"Мой первый пост","content":"Содержание поста"}'

# Получить все посты
curl http://localhost:8080/posts

# Добавить комментарий
curl -X POST http://localhost:8080/posts/1/comments \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -d '{"content":"Отличный пост!"}'

# Получить комментарии
curl http://localhost:8080/posts/1/comments
```

## Задания на неделю

### День 22-23: Базовая структура
- Создайте проект и структуру папок
- Реализуйте модели
- Реализуйте storage

### День 24-25: Handlers
- Реализуйте все CRUD операции для Users
- Реализуйте все CRUD операции для Posts
- Добавьте валидацию

### День 26: Комментарии и связи
- Реализуйте комментарии к постам
- Добавьте получение постов пользователя
- Добавьте пагинацию (limit, offset)

### День 27: Улучшения
- Добавьте поиск по постам
- Добавьте сортировку (по дате, популярности)
- Улучшите обработку ошибок

### День 28: Финал
- Протестируйте все endpoints
- Добавьте README с документацией
- Опубликуйте на GitHub

## Бонус: Что добавить дальше

1. **Аутентификация**: JWT токены
2. **База данных**: PostgreSQL или SQLite
3. **Тесты**: unit и integration тесты
4. **Docker**: контейнеризация
5. **Swagger**: документация API
