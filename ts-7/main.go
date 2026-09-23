package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	numbers := make([]int, 1000000)
	for i := range 1000000 {
		numbers[i] = i
	}

	primeNumbers := fanIn(ctx, fanOut(ctx, generate(ctx, numbers...)))

	for primeNumber := range primeNumbers {
		fmt.Println(primeNumber)
	}

}

func generate(ctx context.Context, numbers ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range numbers {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()

	return out
}

func fanOut[T int](ctx context.Context, in <-chan T) []<-chan T {
	availableCpu := runtime.NumCPU()
	outs := make([]<-chan T, availableCpu)

	for i := range availableCpu {
		outs[i] = primeFinder(ctx, in)
	}

	return outs
}

func primeFinder[T int](ctx context.Context, in <-chan T) <-chan T {

	findPrime := func(value T) bool {
		if value <= 1 {
			return false
		}
		for i := T(2); i*i <= value; i++ {
			if value%i == 0 {
				return false
			}
		}

		return true
	}

	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				if findPrime(val) {
					select {
					case <-ctx.Done():
						return
					case out <- val:
					}
				}
			}
		}

	}()

	return out

}

func fanIn[T any](ctx context.Context, channels []<-chan T) <-chan T {
	wg := sync.WaitGroup{}
	out := make(chan T)

	for _, ch := range channels {
		wg.Add(1)
		go func(channel <-chan T) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-channel:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return
					case out <- val:
					}

				}
			}

		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func pool[T any](ctx context.Context, in <-chan T) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}

	availableCpu := runtime.NumCPU()

	for range availableCpu {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- val:
					}
				}
			}

		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func FanOn(ctx context.Context, chans []<-chan int) <-chan int {

	out := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		defer close(out)

		for _, ch := range chans {
			wg.Add(1)

			go func(channel <-chan int) {
				defer wg.Done()

				for {
					select {
					case <-ctx.Done():
						return
					case val, ok := <-channel:
						if !ok {
							return
						}
						select {
						case <-ctx.Done():
							return
						case out <- val:
						}
					}
				}
			}(ch)
		}
		wg.Wait()
	}()

	return out
}

func FanIn(ctx context.Context, chans []<-chan int) <-chan int {

	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- val:
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func FanOut(ctx context.Context, ch <-chan int) []<-chan int {

	availableCpu := runtime.NumCPU()
	outs := make([]<-chan int, availableCpu)

	for i := range availableCpu {
		outs[i] = primeFinder(ctx, ch)
	}

	return outs
}
