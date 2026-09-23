package main

import (
	"context"
	"fmt"
	"go-initial/ts-8/msg"
	"go-initial/ts-8/wp"
	"time"
)

type IPool interface {
	Submit(ctx context.Context, job msg.Message) bool
	Stop()
}

func processData(workerId int, data msg.Message) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d processing %s\n", workerId, data.Msg)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var p IPool
	p = wp.NewPool(ctx, 10, processData)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			messages := msg.GetMessages()
			if len(messages) == 0 {
				continue
			}
			for _, m := range messages {
				p.Submit(ctx, m)
			}
		}
	}

	p.Stop()

}
