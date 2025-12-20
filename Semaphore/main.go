package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

l:
	for {
		select {
		case <-ctx.Done():
			break l
		case <-ticker.C:
			getSomething()
		}
	}
	fmt.Println("Count", count, "GOOOOO", gocount)

}

var gocount = 0
var count = 0
var mx = &sync.Mutex{}

func getSomething() {

	Semaphore(func() {
		getCash()
	})
	mx.Lock()
	count++
	mx.Unlock()
}

func getCash() {
	time.Sleep(time.Millisecond * 300)
	mx.Lock()
	gocount++
	mx.Unlock()
}

var MaxGorutine = 10
var chMax = make(chan struct{}, MaxGorutine)

func Semaphore(f func()) {
	select {
	case chMax <- struct{}{}:
	default:
		return
	}

	go func() {
		f()

		<-chMax
	}()
}
