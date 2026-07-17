package main

import (
	"context"
	"fmt"
	"go-initial/teechan"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	initialChan := setUpChan(ctx, "El and Mike")
	start := time.Now()
	wg := sync.WaitGroup{}
	treeTeeChannels := teechan.New(3)
	treeChannels := treeTeeChannels.Execute(ctx, initialChan)

	for i, channel := range treeChannels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter := 0
			for {
				if i == counter {
					fmt.Printf("VAL============== %v _______ %v \n", i, counter)
					time.Sleep(time.Second)
					fmt.Println("+++IT WORKS+++")
				}
				select {
				case <-ctx.Done():
					return
				case value, ok := <-channel:
					if !ok {
						return
					}
					fmt.Println(value)
					counter++
				}
			}
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println(duration)
}

func setUpChan(parentCtx context.Context, value string) <-chan string {
	outputChan := make(chan string)

	go func() {
		defer close(outputChan)
		for i := 0; i < 5; i++ {
			select {
			case <-parentCtx.Done():
				return
			case outputChan <- fmt.Sprintf("Hello, %s!, Your number is %v", value, i):
			}
		}
	}()

	return outputChan
}

func standardTee(ctx context.Context, inputChan <-chan string) (chan string, chan string) {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		defer close(out1)
		defer close(out2)
		for val := range inputChan {
			out1, out2 := out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-ctx.Done():
					return
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}

func betterTee(ctx context.Context, inputChan <-chan string) (chan string, chan string) {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		defer close(out1)
		defer close(out2)
		for value := range inputChan {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <-ctx.Done():
					return
				case out1 <- value:
				}
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <-ctx.Done():
					return
				case out2 <- value:
				}
			}()
			wg.Wait()
		}
	}()

	return out1, out2
}

func incredibleTee(ctx context.Context, inputChan <-chan string) {

}

func fanIn(ctx context.Context, channels ...chan string) chan string {

	out := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case value, ok := <-channel:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- value:
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
