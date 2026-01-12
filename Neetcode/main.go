package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [][]int{{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	Binary2dMatrix(arr, 6)
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
