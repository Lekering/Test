package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int, 20) // буферизованный канал на 20 элементов
	for i := range 20 {
		ch <- i
	}
	close(ch)
	wg.Add(2)
	go consumer(wg, ch)
	go consumer(wg, ch)
	wg.Wait()
}

func Revers(array []int) {
	f, s := 0, len(array)-1

	for f < s {
		array[f], array[s] = array[s], array[f]
		f++
		s--
	}
	fmt.Println(array)
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}
