package main

import (
	"context"
	"fmt"
	"go-initial/ts-8/msg"
	"go-initial/ts-8/pool"
	"time"
)

type IPool interface {
	Create()
	Handle(data msg.Message)
	Wait()
	Stats()
}

func processData(workerId int, data msg.Message) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d processing %s\n", workerId, data.Msg)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var p IPool
	p = pool.New(processData)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
		}

		messages := msg.GetMessages()
		p.Create()

		for _, m := range messages {
			p.Handle(m)
		}
		p.Wait()
	}
	p.Stats()
}
