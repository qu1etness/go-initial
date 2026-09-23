package main

import (
	"context"
	"fmt"
	"go-initial/pulik/pool"
	"math/rand"
	"time"
)

type IPool interface {
	Create(ctx context.Context)
	Handle(ctx context.Context, data Message)
	Wait()
	Stats()
}

type Message struct {
	ID  int
	Msg string
}

var index int

func getMessage() []Message {

	maxCapacity := rand.Intn(10)

	out := make([]Message, 0, maxCapacity)

	for range maxCapacity {

		out = append(out, Message{
			ID:  index,
			Msg: fmt.Sprintf("Message-%d", index),
		})
		index++

	}
	return out
}

func processData(workerId int, data Message) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d processing %s\n", workerId, data.Msg)
}

func main() {
	var p IPool
	p = pool.New(processData)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

pool:
	for {
		select {
		case <-ctx.Done():
			break pool
		default:
			messages := getMessage()
			if len(messages) == 0 {
				continue
			}
			p.Create(ctx)
			for _, m := range messages {
				p.Handle(ctx, m)
			}
			p.Wait()
		}
	}
	p.Stats()
}
