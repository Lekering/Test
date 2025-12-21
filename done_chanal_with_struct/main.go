package main

import (
	"fmt"
	"time"
)

type Worker struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewWoeker() Worker {
	worker := Worker{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
	}

	go func() {
		defer close(worker.closeDoneCh)
		for {
			select {
			case <-worker.closeCh:
				return
			default:
				select {
				case <-worker.closeCh:
					return
				case <-time.Tick(time.Second):
					fmt.Println("Tick")
				}
			}
		}
	}()
	return worker
}

func (w Worker) DoneChanal() {
	close(w.closeCh)
	<-w.closeDoneCh
	fmt.Println("Game Over")
}

func main() {
	w := NewWoeker()
	time.Sleep(time.Second * 5)
	w.DoneChanal()

	time.Sleep(time.Second * 2)
	fmt.Println("Main done")
}
