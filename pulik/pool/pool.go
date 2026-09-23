package pool

import (
	"context"
	"fmt"
)

type Worker struct {
	ID         int
	jobCounter int
}

var workers = []*Worker{
	{ID: 1},
	{ID: 2},
	{ID: 3},
	{ID: 4},
	{ID: 5},
	{ID: 6},
	{ID: 7},
	{ID: 8},
	{ID: 9},
	{ID: 10},
}

type Pool[D any] struct {
	pool    chan *Worker
	handler func(id int, data D)
}

func New[D any](handler func(id int, data D)) *Pool[D] {
	return &Pool[D]{pool: make(chan *Worker, len(workers)), handler: handler}
}

func (p *Pool[D]) Create(ctx context.Context) {
	for _, w := range workers {
	pushWorker:
		for {
			select {
			case <-ctx.Done():
				return
			case p.pool <- w:
				break pushWorker
			}
		}
	}
}

func (p *Pool[D]) Handle(ctx context.Context, data D) {
	for {
		select {
		case <-ctx.Done():
			return
		case w := <-p.pool:
			go func() {
				p.handler(w.ID, data)
				w.jobCounter++
				p.pool <- w
			}()
			return
		}
	}
}

func (p *Pool[D]) Wait() {
	for range workers {
		<-p.pool
	}
}

func (p *Pool[D]) Stats() {
	fmt.Printf("________________Stats________________\n")
	for _, w := range workers {
		fmt.Printf("Worker %d has %d jobs\n", w.ID, w.jobCounter)
	}
	fmt.Printf("________________Stats________________\n")
}
