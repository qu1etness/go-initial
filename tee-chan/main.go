package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tee := teeChannel(ctx, 3, generate(ctx, 10))

	wg := sync.WaitGroup{}

	for i, ch := range tee {
		wg.Add(1)
		go func(i int, ch <-chan int) {
			defer wg.Done()
			time.Sleep(time.Second)
			for v := range ch {
				fmt.Println("Channel ", i, " recived: ", v)
			}
		}(i, ch)
	}
	wg.Wait()

}

func generate(ctx context.Context, limit int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for i := 1; i < limit; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()

	return out

}

func teeChannel(ctx context.Context, numChannel int, in <-chan int) []chan int {

	outs := make([]chan int, numChannel)
	for i := range outs {
		outs[i] = make(chan int)
	}

	go func() {
		defer func() {
			for _, ch := range outs {
				close(ch)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				wg := sync.WaitGroup{}
				for _, ch := range outs {
					wg.Add(1)
					go func(ch chan int) {
						defer wg.Done()

						select {
						case <-ctx.Done():
							return
						case ch <- val:
							return
						}

					}(ch)
				}
				wg.Wait()

			}
		}

	}()

	return outs
}
