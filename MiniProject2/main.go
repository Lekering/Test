package main

import "fmt"

type cmd func(in, out chan any)

func Pipeline(in chan any, methods ...cmd) chan any {

	currentChan := in

	for _, comand := range methods {
		out := make(chan any)
		go comand(currentChan, out)
		currentChan = out
	}
	return currentChan
}

func generator() chan any {
	out := make(chan any)

	go func() {
		defer close(out)
		for i := range 10 {
			out <- i
		}
	}()
	return out
}

func ReturnOnlySecondNumb() cmd {
	return func(in, out chan any) {
		defer close(out)

		for val := range in {
			if num, ok := val.(int); ok {
				out <- num * 2
			} else {
				out <- num
			}
		}
	}
}

func ReturnRowOnNumber(number int) cmd {
	return func(in, out chan any) {
		defer close(out)

		for val := range in {
			if num, ok := val.(int); ok {
				out <- num + number
			} else {
				out <- number
			}
		}
	}
}

func main() {

	result := Pipeline(
		generator(),
		ReturnOnlySecondNumb(),
		ReturnRowOnNumber(5),
	)

	for val := range result {
		fmt.Println(val)
	}
}
