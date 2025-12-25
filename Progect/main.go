package main

import (
	"fmt"
	"sort"
	"sync"
)

// ExecutePipeline запускает конвейерную обработку
func ExecutePipeline(freeFlowJobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, j := range freeFlowJobs {
		out := make(chan interface{})
		wg.Add(1)
		go func(jFunc job, input, output chan interface{}) {
			defer wg.Done()
			defer close(output)
			jFunc(input, output)
		}(j, in, out)
		in = out
	}

	wg.Wait()
}

// SelectUsers читает email'ы и возвращает уникальных пользователей
func SelectUsers(in, out chan any) {
	wg := &sync.WaitGroup{}
	seenUsers := &sync.Map{}

	for email := range in {
		emailStr := email.(string)
		wg.Add(1)

		go func(e string) {
			defer wg.Done()
			user := GetUser(e)

			// Проверяем, видели ли мы этого пользователя
			if _, loaded := seenUsers.LoadOrStore(user.ID, true); !loaded {
				out <- user
			}
		}(emailStr)
	}

	wg.Wait()
}

// SelectMessages получает сообщения пользователей батчами
func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	batchSize := 2
	batch := make([]User, 0, batchSize)
	mu := &sync.Mutex{}

	processBatch := func(users []User) {
		wg.Add(1)
		go func(u []User) {
			defer wg.Done()
			messages := GetMessages(u...)
			for _, msgID := range messages {
				out <- msgID
			}
		}(users)
	}

	for user := range in {
		u := user.(User)
		mu.Lock()
		batch = append(batch, u)

		if len(batch) >= batchSize {
			processBatch(batch)
			batch = make([]User, 0, batchSize)
		}
		mu.Unlock()
	}

	// Обработать оставшиеся элементы
	if len(batch) > 0 {
		processBatch(batch)
	}

	wg.Wait()
}

// CheckSpam проверяет сообщения на спам с ограничением параллелизма
func CheckSpam(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	semaphore := make(chan struct{}, 5) // Ограничение в 5 параллельных запросов

	for msgID := range in {
		msg := msgID.(MsgID)
		wg.Add(1)

		go func(m MsgID) {
			defer wg.Done()

			semaphore <- struct{}{} // Захватываем слот
			hasSpam := HasSpam(m)
			<-semaphore // Освобождаем слот

			out <- MsgData{
				ID:      m,
				HasSpam: hasSpam,
			}
		}(msg)
	}

	wg.Wait()
}

// CombineResults собирает и сортирует результаты
func CombineResults(in, out chan interface{}) {
	results := make([]MsgData, 0)

	for data := range in {
		msgData := data.(MsgData)
		results = append(results, msgData)
	}

	// Сортировка: сначала по HasSpam (true первые), затем по ID
	sort.Slice(results, func(i, j int) bool {
		if results[i].HasSpam != results[j].HasSpam {
			return results[i].HasSpam
		}
		return results[i].ID < results[j].ID
	})

	for _, result := range results {
		out <- fmt.Sprintf("%t %d", result.HasSpam, result.ID)
	}
}

// Типы данных для примера (нужно заменить на реальные)
type job func(in, out chan interface{})

type User struct {
	ID    int
	Email string
}

type MsgID int

type MsgData struct {
	ID      MsgID
	HasSpam bool
}

// Пример запуска
func main() {
	// Входные данные - email'ы пользователей
	inputData := []string{
		"user1@example.com",
		"user2@example.com",
		"batman@mail.ru",
		"bruce.wayne@mail.ru", // алиас к batman
		"user3@example.com",
	}

	// Создаем генератор входных данных
	dataSource := func(in, out chan interface{}) {
		for _, email := range inputData {
			out <- email
		}
	}

	// Создаем приемник результатов
	dataSink := func(in, out chan interface{}) {
		for result := range in {
			fmt.Println(result)
		}
	}

	// Запускаем pipeline
	ExecutePipeline(
		job(dataSource),
		job(SelectUsers),
		job(SelectMessages),
		job(CheckSpam),
		job(CombineResults),
		job(dataSink),
	)
}

// Заглушки для функций API (замените на реальные)
func GetUser(email string) User {
	// Симуляция задержки
	return User{ID: 1, Email: email}
}

func GetMessages(users ...User) []MsgID {
	// Симуляция задержки
	return []MsgID{1, 2, 3}
}

func HasSpam(msgID MsgID) bool {
	// Симуляция задержки
	return false
}
