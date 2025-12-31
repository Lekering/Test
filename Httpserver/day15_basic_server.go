package main

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик для главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет! Это мой первый HTTP сервер на Go!")
}

// Обработчик для страницы "о нас"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Страница 'О нас'\nМетод запроса: %s", r.Method)
}

// Обработчик с параметрами из URL
func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр из URL: /greet?name=Вася
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "гость"
	}
	fmt.Fprintf(w, "Привет, %s!", name)
}

func main() {
	// Регистрируем обработчики для разных путей
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/greet", greetHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
ЗАДАНИЯ НА ДЕНЬ 15:
1. Запустите: go run main.go
2. Откройте в браузере:
   - http://localhost:8080/
   - http://localhost:8080/about
   - http://localhost:8080/greet?name=Иван

3. ПРАКТИКА:
   - Добавьте обработчик /time, который возвращает текущее время
   - Добавьте обработчик /sum?a=5&b=3, который суммирует два числа
   - Попробуйте добавить HTML в ответ
*/
