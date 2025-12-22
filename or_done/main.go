package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan string)

	go func() {
		defer close(in)
		for range 10 {
			time.Sleep(time.Millisecond * 100)
			in <- "Hi"
		}
	}()

	doneCh := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(doneCh)
	}()

	for value := range OrDone(in, doneCh) {
		fmt.Println(value)
	}

}

func OrDone[T any](in chan T, doneCh chan struct{}) chan T {
	outCh := make(chan T)

	go func() {
		defer close(outCh)

		for {
			select {
			case <-doneCh:
				fmt.Println("End1")
				return
			default:
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					time.Sleep(time.Second)
					outCh <- v
				case <-doneCh:
					fmt.Println("End2")
					return
				}
			}
		}
	}()

	return outCh
}
