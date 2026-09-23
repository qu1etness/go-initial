package main

import (
	"context"
	"sync"
)

type Pool[T any] struct {
	jobs       chan T
	maxWorkers int
	wg         sync.WaitGroup
	handler    func(workerID int, job T)
}

func New[T any](ctx context.Context, handler func(workerID int, job T), maxWorkers int) *Pool[T] {

	p := &Pool[T]{
		jobs:       make(chan T),
		maxWorkers: maxWorkers,
		wg:         sync.WaitGroup{},
	}

	for id := 0; id < maxWorkers; id++ {
		p.wg.Add(1)
		go func(wID int) {
			defer p.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-p.jobs:
					if !ok {
						return
					}
					p.handler(wID, job)
				}
			}

		}(id)
	}

	return p

}

func (p *Pool[T]) Submit(ctx context.Context, job T) bool {
	select {
	case <-ctx.Done():
		return false
	case p.jobs <- job:
		return true
	}
}

func (p *Pool[T]) Stop() {
	close(p.jobs)
	p.wg.Wait()
}
