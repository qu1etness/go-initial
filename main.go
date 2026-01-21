package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{})

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-done:
				return
			case cows <- "Myyyy":
			case pigs <- "Ionk":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows)
	wg.Add(1)
	go consumePigs(done, pigs)

	wg.Wait()
}

func consumeCows(done, cows chan interface{}) {
	defer wg.Done()
	for result := range orDone(done, cows) {
		fmt.Println(result)
	}
}

func consumePigs(done, pigs chan interface{}) {
	defer wg.Done()
	for result := range orDone(done, pigs) {
		fmt.Println(result)
	}
}

func orDone(done, inputChan <-chan interface{}) <-chan interface{} {
	outputChan := make(chan interface{})

	go func() {
		defer close(outputChan)
		for {
			select {
			case <-done:
				return
			case value, isOk := <-inputChan:
				if !isOk {
					return
				}
				select {
				case outputChan <- value:
				case <-done:
					return
				}
			}

		}
	}()

	return outputChan

}
