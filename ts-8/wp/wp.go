package wp

import (
	"context"
	"sync"
)

type WorkerPool[T any] struct {
	jobs    chan T
	workers int
	wg      sync.WaitGroup
}

func NewPool[T any](ctx context.Context, workerCount int, handler func(workerID int, job T)) *WorkerPool[T] {
	p := &WorkerPool[T]{
		jobs:    make(chan T),
		workers: workerCount,
	}
	for id := 1; id <= workerCount; id++ {
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
					handler(wID, job)
				}
			}

		}(id)
	}
	return p
}

func (p *WorkerPool[T]) Submit(ctx context.Context, job T) bool {
	select {
	case <-ctx.Done():
		return false
	case p.jobs <- job:
		return true
	}
}

func (p *WorkerPool[T]) Stop() {
	close(p.jobs)
	p.wg.Wait()
}
