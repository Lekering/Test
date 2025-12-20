package main

import "fmt"

type Worker struct {
	id       int
	countJob int
}

var worker = []Worker{
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
}

type Pool[Data any] struct {
	pool    chan *Worker
	handler func(int, Data)
}

func NewPool[Data any](hadle func(int, Data)) *Pool[Data] {
	return &Pool[Data]{
		pool:    make(chan *Worker, len(worker)),
		handler: hadle,
	}
}

func (p *Pool[Data]) Create() {
	for i := range worker {
		p.pool <- &worker[i]
	}
}

func (p *Pool[Data]) Handle(d Data) {
	w := <-p.pool
	go func() {
		p.handler(w.id, d)
		w.countJob++
		p.pool <- w
	}()
}

func (p *Pool[Data]) Wait() {
	for range len(worker) {
		<-p.pool
	}
}

func (p *Pool[Data]) Info() {
	fmt.Println("____________________RESULT______________________")
	for _, w := range worker {
		fmt.Printf("Worker Number: %v completed %v jobs\n", w.id, w.countJob)
	}
}
