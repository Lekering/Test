package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	chan1in, chan2in, chan3in := make(chan int), make(chan int), make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	go func() {
		defer close(chan1in)
		defer close(chan2in)
		defer close(chan3in)

		for i := range 1000 {
			time.Sleep(time.Second)
			chan1in <- i
			chan2in <- i + 10
			chan3in <- i + 100
		}
	}()

	for value := range MergeChanel(ctx, chan1in, chan2in, chan3in) {
		fmt.Println(value)
	}
}

func MergeChanel[T any](ctx context.Context, inValueChan ...<-chan T) <-chan T {
	outChan := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(inValueChan))

	for _, inChan := range inValueChan {
		go func() {
			defer wg.Done()
			for {
				select {
				case v, ok := <-inChan:
					if !ok {
						return
					}
					select {
					case outChan <- v:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	return outChan
}
