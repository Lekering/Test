package main

import (
	"fmt"
	"log"
)

func main() {
	array := []int{4, 3, 2, 1}
	fmt.Printf("array: %v\n", array)
	Revers(array)
	fmt.Printf("array: %v\n", array)
}

func Revers(arr []int) {
	f, s := 0, len(arr)-1

	for f < s {
		arr[f], arr[s] = arr[s], arr[f]
		f++
		s--
		log.Print(arr)
	}
}
