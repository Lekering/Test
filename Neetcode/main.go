package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("generateParenthesis(3): %v\n", generateParenthesis(3))
}

func Revers(array []int) {
	f, s := 0, len(array)-1

	for f < s {
		array[f], array[s] = array[s], array[f]
		f++
		s--
	}
	fmt.Println(array)
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}
