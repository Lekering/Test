package main

import (
	"fmt"
	"log"
)

func main() {
	array := []int{67, 6, 11, 2, 9, 25, 1, 3, 110, 5, 4, 10} // элементы перемешаны
	//fmt.Printf("BubleSort(array): %v\n", BubleSort(array))
	fmt.Printf("QuickSort(array): %v\n", QuickSort(array))
	fmt.Printf("BinarySearch(6, array...): %v\n", BinarySearch(11, array))
	array = append(array, array...)
}

func BinarySearch(target int, array []int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		midle := (left + right) / 2
		if target == array[midle] {
			return midle
		} else if array[midle] < target {
			left = midle + 1
		} else if array[midle] > target {
			right = midle
		}
	}
	return -1
}

func BubleSort(array []int) []int {
	n := len(array) - 1

	for i := range n {
		for j := 0; j < n-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
		log.Print(array)
	}
	return array
}

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left := 0
	right := len(array) - 1
	pivotIndex := (left + right) / 2
	pivot := array[pivotIndex]

	for left <= right {
		for array[left] < pivot {
			left++
		}
		for array[right] > pivot {
			right--
		}
		if left <= right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
		log.Print(array)
	}
	if right > 0 {
		QuickSort(array[:right+1])
	}
	if left < len(array) {
		QuickSort(array[left:])
	}
	return array
}
