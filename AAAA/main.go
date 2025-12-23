package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main_() {
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}
	counter := 1000

	doublec := make([]int, 0, counter)
	limitedNum := make(chan int)
	counterMap := make(map[int]struct{})

	for range counter {
		doublec = append(doublec, rand.Intn(10))
	}

	for i := range counter {
		wg.Add(1)
		go func(int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := counterMap[doublec[i]]; !ok {
				counterMap[doublec[i]] = struct{}{}
				limitedNum <- doublec[i]
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(limitedNum)
	}()

	for v := range limitedNum {
		fmt.Println(v)
	}

}
