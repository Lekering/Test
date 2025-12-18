package main

import "fmt"

type Worker struct {
	id   int
	jobs int
}

var workers = []Worker{
	{id: 1},
	{id: 2},
	{id: 3},
	{id: 4},
	{id: 5},
	{id: 6},
	{id: 7},
	{id: 8},
	{id: 9},
	{id: 10},
	{id: 11},
}

type Pool[Data any] struct {
	pool    chan *Worker
	handler func(int, Data)
}

func New[Data any](handler func(int, Data)) *Pool[Data] {
	return &Pool[Data]{
		handler: handler,
		pool:    make(chan *Worker, len(workers)),
	}
}

func (p *Pool[Data]) Create() {
	for i := range workers {
		// кладём в канал указатель на реальный элемент слайса, а не на копию
		p.pool <- &workers[i]
	}

}

func (p *Pool[Data]) Handle(d Data) {
	w := <-p.pool
	go func() {
		p.handler(w.id, d)
		w.jobs++
		p.pool <- w
	}()
}

func (p *Pool[Data]) Wait() {
	for range len(workers) {
		<-p.pool
	}
}

func (p *Pool[Data]) Statistk() {
	fmt.Println("_________________________Resalts___________________________")
	for _, w := range workers {
		fmt.Printf("Worker Number: %d Job Complited: %d\n", w.id, w.jobs)
	}
}
