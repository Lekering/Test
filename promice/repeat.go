package main

type Result[T any] struct {
	value T
	err   error
}

type ChanRes[T any] struct {
	resChan chan Result[T]
}

func NewChan[T any](f func() (T, error)) *ChanRes[T] {
	res := ChanRes[T]{
		resChan: make(chan Result[T]),
	}

	go func() {
		defer close(res.resChan)

		val, err := f()
		res.resChan <- Result[T]{value: val, err: err}
	}()

	return &res
}

func (c *ChanRes[T]) ThenRepeat(t func(T), f func(error)) {
	go func() {
		res := <-c.resChan
		if res.err == nil {
			t(res.value)
		} else {
			f(res.err)
		}
	}()
}
