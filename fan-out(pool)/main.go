package main

import (
	"fmt"
	"sync"
	"time"
)

func generate() chan int {
	in := make(chan int)
	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()
	return in
}

func main() {
	for v := range Pool(generate(), 100, Math) {
		fmt.Println(v)
	}
}

func Pool[T any](in chan T, numberWorker int, f func(T) T) <-chan T {
	outCh := make(chan T)
	wg := &sync.WaitGroup{}
	go func() {
		for range numberWorker {
			wg.Add(1)
			go Work(in, outCh, wg, f)
		}
		// Ждем завершения всех горутин перед закрытием канала
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func Work[T any](in, out chan T, wg *sync.WaitGroup, f func(T) T) {
	defer wg.Done()

	for v := range in {
		out <- f(v)
	}
}

func Math(n int) int {
	time.Sleep(time.Millisecond * 1000)
	n *= n
	return n
}
