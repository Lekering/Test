package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ============= MIDDLEWARE =============

// Логирование запросов
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Создаем обертку для ResponseWriter, чтобы захватить статус код
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(lrw, r)
		
		duration := time.Since(start)
		log.Printf("[%s] %s %s - %d (%v)",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			lrw.statusCode,
			duration,
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// CORS middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Обрабатываем preflight запросы
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// Простая аутентификация через токен
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		
		// Простая проверка токена (в реальности - JWT или OAuth)
		if token != "Bearer secret-token-123" {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		
		// Добавляем user ID в контекст
		ctx := context.WithValue(r.Context(), "userID", 42)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Rate limiting (упрощенный)
type RateLimiter struct {
	requests map[string][]time.Time
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		now := time.Now()
		limit := 10 // 10 запросов
		window := time.Minute
		
		// Очищаем старые запросы
		if times, exists := rl.requests[ip]; exists {
			var recent []time.Time
			for _, t := range times {
				if now.Sub(t) < window {
					recent = append(recent, t)
				}
			}
			rl.requests[ip] = recent
		}
		
		// Проверяем лимит
		if len(rl.requests[ip]) >= limit {
			http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
			return
		}
		
		// Добавляем текущий запрос
		rl.requests[ip] = append(rl.requests[ip], now)
		
		next.ServeHTTP(w, r)
	})
}

// Recovery middleware (ловит панику)
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)
				http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// ============= HANDLERS =============

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "API работает!",
		"version": "1.0",
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем userID из контекста (добавлен в AuthMiddleware)
	userID := r.Context().Value("userID").(int)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Это защищенный endpoint",
		"user_id": userID,
	})
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	// Эта паника будет поймана RecoveryMiddleware
	panic("Тестовая паника!")
}

// ============= MAIN =============

func main() {
	mux := http.NewServeMux()
	
	// Публичные endpoints
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/panic", panicHandler)
	
	// Защищенные endpoints (требуют авторизацию)
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/protected", protectedHandler)
	
	// Оборачиваем защищенные routes в AuthMiddleware
	mux.Handle("/protected", AuthMiddleware(protectedMux))
	
	// Создаем rate limiter
	rateLimiter := NewRateLimiter()
	
	// Применяем middleware в порядке (снизу вверх):
	// 1. Recovery (самый внешний - ловит все панику)
	// 2. Logging (логирует все запросы)
	// 3. CORS (добавляет CORS заголовки)
	// 4. RateLimit (ограничивает частоту запросов)
	// 5. mux (роутер с handlers)
	handler := RecoveryMiddleware(
		LoggingMiddleware(
			CORSMiddleware(
				rateLimiter.Middleware(mux),
			),
		),
	)
	
	fmt.Println("🚀 Сервер с middleware запущен на :8080")
	fmt.Println("\nEndpoints:")
	fmt.Println("  GET  / - публичный")
	fmt.Println("  GET  /protected - требует токен")
	fmt.Println("  GET  /panic - тест recovery")
	fmt.Println("\nТестирование:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/protected")
	fmt.Println(`  curl -H "Authorization: Bearer secret-token-123" http://localhost:8080/protected`)
	
	log.Fatal(http.ListenAndServe(":8080", handler))
}

/*
ЗАДАНИЯ НА ДЕНЬ 20-21:

1. Запустите и протестируйте:
   # Публичный endpoint
   curl http://localhost:8080/
   
   # Защищенный (без токена - ошибка)
   curl http://localhost:8080/protected
   
   # Защищенный (с токеном - работает)
   curl -H "Authorization: Bearer secret-token-123" \
        http://localhost:8080/protected
   
   # Тест recovery
   curl http://localhost:8080/panic
   
   # Тест rate limit (сделайте 11+ запросов подряд)
   for i in {1..12}; do curl http://localhost:8080/; done

2. ПРАКТИКА:
   - Добавьте middleware для добавления request ID
   - Реализуйте более продвинутый rate limiter (по токену)
   - Добавьте middleware для валидации Content-Type
   - Создайте middleware для логирования тела запроса

3. ДОПОЛНИТЕЛЬНО:
   - Изучите библиотеку gorilla/mux для роутинга
   - Посмотрите на chi router (более легкий)
   - Почитайте про JWT токены
*/
