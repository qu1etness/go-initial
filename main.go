package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	fruits := []string{"Apple", "Orange", "Peach"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resultArr := make([]chan string, 0, len(fruits))

	for _, fruit := range fruits {
		resultArr = append(resultArr, setToChan(ctx, fruit))
	}

	wg.Add(3)
	go appleFunc(ctx, resultArr[0], &wg)

	otherFruits(ctx, resultArr[1], &wg)

	otherFruits(ctx, resultArr[2], &wg)

	wg.Wait()
}

func setToChan(ctx context.Context, name string) chan string {
	outputChan := make(chan string)

	go func() {
		defer close(outputChan)
		for {
			select {
			case <-ctx.Done():
				return
			case outputChan <- name:
			}
		}
	}()

	return outputChan
}

func appleFunc(parentCtx context.Context, stream <-chan string, parentWg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	wg := sync.WaitGroup{}
	defer parentWg.Done()
	defer cancel()
	ticker := time.NewTicker(500 * time.Millisecond)

	doWork := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					return
				case apple := <-stream:
					fmt.Println(apple, "+++++++++++")
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doWork(ctx)
	}
	wg.Wait()
}

func otherFruits(ctx context.Context, stream <-chan string, parentWg *sync.WaitGroup) {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		defer parentWg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					return
				case fruit := <-stream:
					fmt.Println("=============", fruit)
				}
			}
		}
	}()
}
