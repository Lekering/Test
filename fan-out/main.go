package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	wg := &sync.WaitGroup{}

	go func() {
		defer close(in)
		for i := range 10 {
			in <- i
		}
	}()

	outChans := SplitChan(in, 2)

	for _, out := range outChans {

		wg.Add(1)

		go func(out <-chan int) {
			defer wg.Done()
			for v := range out {
				fmt.Println(v)
			}
		}(out)
	}

	wg.Wait()
}

func SplitChan[T any](in <-chan T, n int) []chan T {

	outChans := make([]chan T, n)
	for i := range outChans {
		outChans[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, c := range outChans {
				close(c)
			}
		}()

		ind := 0
		for value := range in {
			outChans[ind] <- value
			ind = (ind + 1) % n
		}
	}()

	return outChans
}
