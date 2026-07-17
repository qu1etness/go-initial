package teechan

import (
	"context"
	"fmt"
	"sync"
)

type TeeChan struct {
	channels []chan string
	numChan  int
	wgs      []*sync.WaitGroup
}

func New(numChan int) *TeeChan {
	channels := make([]chan string, numChan)
	wgs := make([]*sync.WaitGroup, numChan)

	for i := range numChan {
		channels[i] = make(chan string)
		wgs[i] = &sync.WaitGroup{}
	}

	return &TeeChan{
		channels,
		numChan,
		wgs,
	}
}

func (t *TeeChan) Execute(ctx context.Context, inputChan <-chan string) []chan string {
	go func() {

		defer func() {
			for i := range t.numChan {
				go func() {
					defer close(t.channels[i])
					t.wgs[i].Wait()
					fmt.Println("Everything is don")
				}()
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case inputValue, ok := <-inputChan:
				if !ok {
					return
				}
				wg := sync.WaitGroup{}

				for i := range t.numChan {
					//t.wgs[i].Add(1)
					wg.Add(1)
					go func() {
						defer wg.Done()
						//defer t.wgs[i].Done()
						select {
						case <-ctx.Done():
							return
						case t.channels[i] <- inputValue:
						}
					}()
				}
				wg.Wait()
			}
		}
	}()

	return t.channels

}
