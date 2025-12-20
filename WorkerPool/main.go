package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var maxMessagers = 5
var counter = 0

type Message struct {
	id    int
	title string
	text  string
}

type IPool[Data any] interface {
	Create()
	Handle(m Data)
	Wait()
	Statistk()
}

func main() {
	var pool IPool[Message]
	pool = New(processWork)

	ctx, cansel := context.WithTimeout(context.Background(), time.Second*5)
	defer cansel()

l:
	for {
		select {
		case <-ctx.Done():
			break l
		default:
		}
		messages := getMessages()

		pool.Create()
		for _, message := range messages {
			pool.Handle(message)
		}
		pool.Wait()
	}
	pool.Statistk()
}

func processWork(workerId int, mes Message) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Worker number %v, messages %v title %v \n", workerId, mes.id, mes.title)
}

func getMessages() []Message {
	messagerCount := rand.Intn(maxMessagers)

	arrayMessage := make([]Message, 0, messagerCount)

	for range messagerCount {
		counter++
		arrayMessage = append(arrayMessage, Message{id: counter, title: "WOOOOORLD"})
	}
	return arrayMessage
}
