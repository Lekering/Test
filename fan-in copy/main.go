package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	chan1in, chan2in, chan3in := make(chan int), make(chan int), make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
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
	g, ectx := errgroup.WithContext(ctx)

	for _, inChan := range inValueChan {
		// захватываем текущий inChan
		ch := inChan
		g.Go(func() error {
			for {
				select {
				case <-ectx.Done():
					return ectx.Err()
				case v, ok := <-ch:
					if !ok {
						return nil
					}
					// Искусственно бросаем ошибку для проверки errgroup!
					if i, ok := any(v).(int); ok && i == 5 {
						return fmt.Errorf("errgroup test: value is 5")
					}
					select {
					case outChan <- v:
					case <-ectx.Done():
						return ectx.Err()
					}
				}
			}
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			log.Printf("caught errgroup error: %v", err)
		}
		close(outChan)
	}()

	return outChan
}
