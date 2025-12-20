package main

import "fmt"

func main() {
	ch := make(chan int)

	fn := func(n int) int {
		return n * n
	}

	go func() {
		defer close(ch)
		for i := range 20 {
			ch <- i
		}
	}()

	for v := range Transformer(ch, fn) {
		fmt.Println(v)
	}

}

func Transformer(in <-chan int, f func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- f(v)
		}
	}()
	return out
}
