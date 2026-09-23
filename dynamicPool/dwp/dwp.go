package dwp

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type DynamicPool[T any] struct {
	minWorkers  int
	maxWorkers  int
	idleTimeout time.Duration

	jobs           chan T
	currentWorkers int32
	nextWorkerID   int32

	handler func(workerID int, job T)
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func New[T any](minWorkers int, maxWorkers int, idleTimeout time.Duration, handler func(workerID int, job T)) *DynamicPool[T] {
	dp := &DynamicPool[T]{
		minWorkers:  minWorkers,
		maxWorkers:  maxWorkers,
		idleTimeout: idleTimeout,
		jobs:        make(chan T),
		handler:     handler,
	}

	for range minWorkers {
		dp.spawnPool(true)
	}

	return dp
}

func (dp *DynamicPool[T]) spawnPool(isCore bool) bool {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	if dp.currentWorkers >= int32(dp.maxWorkers) {
		return false
	}

	dp.currentWorkers++
	dp.nextWorkerID++
	workerID := dp.nextWorkerID

	dp.wg.Add(1)
	go func(workerID int, isCore bool) {
		defer func() {
			atomic.AddInt32(&dp.currentWorkers, -1)
			dp.wg.Done()
		}()

		ticker := time.NewTicker(dp.idleTimeout)
		defer ticker.Stop()

		for {
			if isCore {
				job, ok := <-dp.jobs
				if !ok {
					return
				}
				dp.handler(workerID, job)

			} else {
				select {
				case <-ticker.C:
					fmt.Printf("📉 [Scale Down] Worker %d idle for %v -> Exiting\n", workerID, dp.idleTimeout)
					return
				case job, ok := <-dp.jobs:
					if !ok {
						return
					}
					dp.handler(workerID, job)
				}
			}
		}

	}(int(workerID), isCore)

	return true
}

func (dp *DynamicPool[T]) Submit(ctx context.Context, job T) {
	select {
	case <-ctx.Done():
		return
	case dp.jobs <- job:
		return
	default:
		dp.spawnPool(false)
		for {
			select {
			case <-ctx.Done():
				return
			case dp.jobs <- job:
				return
			}
		}
	}
}

func (p *DynamicPool[T]) Stop() {
	close(p.jobs)
	p.wg.Wait()
}

func (p *DynamicPool[T]) ActiveWorkersCount() int {
	return int(atomic.LoadInt32(&p.currentWorkers))
}

func main() {

}
