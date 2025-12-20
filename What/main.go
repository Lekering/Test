package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Воркер %d начал задачу %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Запускаем 3 воркера
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Отправляем 9 задач
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Получаем результаты
	for r := 1; r <= 9; r++ {
		<-results
	}
}
