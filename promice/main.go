package main

import (
	"fmt"
	"time"
)

type result[T any] struct {
	value T
	err   error
}

type Promice[T any] struct {
	resultCh chan result[T]
}

func NewProm[T any](f func() (T, error)) *Promice[T] {
	prom := &Promice[T]{
		resultCh: make(chan result[T]),
	}

	go func() {
		defer close(prom.resultCh)

		val, err := f()
		prom.resultCh <- result[T]{value: val, err: err}

	}()

	return prom
}

func (p *Promice[T]) Then(true func(T), false func(error)) {
	go func() {
		result := <-p.resultCh
		if result.err == nil {
			true(result.value)
		} else {
			false(result.err)
		}
	}()
}

func main() {
	EasyFunc := func() (int, error) {
		time.Sleep(time.Second)
		return 5, nil
	}

	prom := NewProm(EasyFunc)

	chanWait := make(chan struct{})
	prom.Then(
		func(i int) {
			fmt.Println("Number")
			close(chanWait)
		},
		func(err error) {
			fmt.Println(err.Error())
			close(chanWait)
		},
	)
	<-chanWait
}
