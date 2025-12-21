package main

import (
	"fmt"
	"time"
)

func main() {

	stopCh := make(chan struct{})

	doneCh := Process(stopCh)

	go func() {
		time.Sleep(3 * time.Second)
		close(stopCh)
	}()

	<-doneCh
	fmt.Println("Done")
}

func Process(stopCh chan struct{}) chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)

		for {
			select {
			case <-stopCh:
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("USSSSSA")
			}
		}
	}()

	return ch
}
