package main

import (
	"fmt"
	"sync"
	"time"
)

func generateCh() chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for i := range 100 {
			out <- i
		}
	}()
	return out
}

func parse(in chan int) chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)
		for v := range in {
			if v%2 == 0 && v > 50 {
				outCh <- v
			}
		}
	}()

	return outCh

}

func send(in chan int, numChOut int) chan int {

	wg := &sync.WaitGroup{}
	outCh := make(chan int)
	chSplit := split(in, numChOut)

	for i := range numChOut {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for value := range chSplit[idx] {
				outCh <- value
			}
		}(i)

	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh

}

func split(in chan int, n int) []chan int {

	outCh := make([]chan int, n)

	for ch := range n {
		outCh[ch] = make(chan int)
	}

	go func() {
		defer func() {
			for _, v := range outCh {
				close(v)
			}
		}()

		ind := 0
		for value := range in {
			outCh[ind] <- value
			ind = (ind + 1) % n
		}

	}()

	return outCh
}

func main() {

	for value := range send(parse(generateCh()), 2) {
		time.Sleep(time.Second)
		fmt.Println(value)
	}

}
