package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Massage struct {
	id    int
	title string
	text  string
}

type IPool[T any] interface {
	Create()
	Handle(data T)
	Wait()
	Info()
}

func main() {
	var pool IPool[Massage]
	pool = NewPool(ProcessFunc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

l:
	for {
		select {
		case <-ctx.Done():
			break l
		default:
		}

		dataMassage := TakeMassage(25)
		pool.Create()

		for _, massage := range dataMassage {
			pool.Handle(massage)
		}
		pool.Wait()
	}
	pool.Info()
}

var id = 0

func TakeMassage(count int) []Massage {

	ct := rand.Intn(count)
	result := make([]Massage, 0, ct)

	for range ct {
		id++
		result = append(result, Massage{id: id})
	}
	return result
}

func ProcessFunc(workerId int, massage Massage) {
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("Worker number: %v       massageId: %v\n", workerId, massage.id)
}
