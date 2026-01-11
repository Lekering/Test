package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := []int{-8, -6, 0, 2, 4, 5, 12}
	fmt.Printf("squares(arr): %v\n", Squares(arr))

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
