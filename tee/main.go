package main

import (
	"fmt"
	"sync"
	"time"
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

	for i, out := range outChans {
		wg.Add(1)
		sleepDuration := time.Millisecond * 200
		if i == 1 {
			sleepDuration = time.Millisecond * 100
		}
		go func(out <-chan int, sleep time.Duration) {
			defer wg.Done()
			for v := range out {
				time.Sleep(sleep)
				fmt.Println(v)
			}
		}(out, sleepDuration)
	}

	wg.Wait()
}

func SplitChan[T any](in <-chan T, n int) []chan T {
	wg := &sync.WaitGroup{}

	outChans := make([]chan T, n)
	for i := range outChans {
		outChans[i] = make(chan T)
	}

	go func() {
		defer func() {
			wg.Wait() // ждем завершения всех горутин перед закрытием каналов
			for _, c := range outChans {
				close(c)
			}
		}()
		for v := range in {
			for _, c := range outChans {
				select {
				case c <- v:
				default: // prevent goroutine block if receiver is slow
					wg.Add(1)
					go func(val T, ch chan T) {
						defer wg.Done()
						ch <- val
					}(v, c)
				}
			}
		}
	}()

	return outChans
}
