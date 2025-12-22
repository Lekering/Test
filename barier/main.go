package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	mtx   sync.Mutex
	size  int
	count int

	beforeCh chan struct{}
	afterCh  chan struct{}
}

func NewBarrier(num int) *Barrier {
	return &Barrier{
		size:     num,
		beforeCh: make(chan struct{}, num),
		afterCh:  make(chan struct{}, num),
	}
}

func (b *Barrier) Before() {
	b.mtx.Lock()

	b.count++
	if b.count == b.size {
		for i := 0; i < b.size; i++ {
			b.beforeCh <- struct{}{}
		}
	}
	b.mtx.Unlock()
	<-b.beforeCh
}

func (b *Barrier) After() {
	b.mtx.Lock()

	b.count--
	if b.count == 0 {
		for i := 0; i < b.size; i++ {
			b.afterCh <- struct{}{}
		}
	}
	b.mtx.Unlock()
	<-b.afterCh
}

func main() {
	var wg = sync.WaitGroup{}

	boot := func() {
		fmt.Println("boot")
		time.Sleep(time.Second)
	}

	work := func() {
		fmt.Println("work")
		time.Sleep(time.Second)
	}
	count := 2
	wg.Add(count)
	barrier := NewBarrier(count)
	for range count {
		go func() {
			defer wg.Done()
			for range count {
				barrier.Before()
				boot()
				barrier.After()
				work()
			}
		}()
	}
	wg.Wait()
}
