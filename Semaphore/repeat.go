package main

import "fmt"

var maxGo = make(chan struct{}, 10)

func SemaphoreRepeat(f func()) {
	select {
	case maxGo <- struct{}{}:
	default:
		return
	}

	go func() {
		f()
		<-maxGo
	}()
}

func GetCashe() {
	fmt.Println("LOL")
}

func GetSome() {

	SemaphoreRepeat(func() {
		GetCashe()
	})

	fmt.Println("KEK")
}
