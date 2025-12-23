package main

import "sync"

type Bar struct {
	rm         sync.Mutex
	size       int
	count      int
	chanBefore chan struct{}
	chanAfter  chan struct{}
}

func NewBar(size int) *Bar {
	return &Bar{
		size:       size,
		chanBefore: make(chan struct{}),
		chanAfter:  make(chan struct{}),
	}
}

func (b *Bar) RBefor() {
	b.rm.Lock()

	b.count++
	if b.count == b.size {
		for i := 0; i < b.size; i++ {
			b.chanBefore <- struct{}{}
		}
	}
	b.rm.Unlock()
	<-b.chanBefore
}

func (b *Bar) RAfter() {
	b.rm.Lock()

	b.count--
	if b.count == 0 {
		for i := 0; i < b.size; i++ {
			b.chanAfter <- struct{}{}
		}
	}
	b.rm.Unlock()
	<-b.chanAfter
}
