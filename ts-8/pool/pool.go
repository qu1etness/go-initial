package pool

import "fmt"

type Workers struct {
	ID         int
	JobCounter int
}

var workers = []*Workers{
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
	pool    chan *Workers
	handler func(int, D)
}

func New[D any](handler func(int, D)) *Pool[D] {
	return &Pool[D]{handler: handler, pool: make(chan *Workers, len(workers))}
}

func (p *Pool[D]) Create() {
	for _, w := range workers {
		p.pool <- w
	}
}

func (p *Pool[D]) Handle(data D) {
	w := <-p.pool
	go func() {
		p.handler(w.ID, data)
		w.JobCounter++
		p.pool <- w
	}()
}

func (p *Pool[D]) Wait() {
	for range workers {
		<-p.pool
	}
}

func (p *Pool[D]) Stats() {
	fmt.Printf("________________Stats________________\n")
	for _, w := range workers {
		fmt.Printf("Worker %d has %d jobs\n", w.ID, w.JobCounter)
	}
	fmt.Printf("________________Stats________________\n")
}
