package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Number interface {
	int | int32 | int64
}

func main() {
	start := time.Now()
	done := make(chan bool)
	defer close(done)

	randIntFunc := func() int { return rand.Intn(500000000) }
	stream := repeatAnalog(done, randIntFunc)
	routines := fanOut(done, stream)
	fannedInStream := fanIn(done, routines...)

	for limitedStream := range takeAnalog(done, fannedInStream, 30) {
		fmt.Println(limitedStream)
	}

	duration := time.Since(start)
	fmt.Println(duration)
}

func fanIn[T Number, K bool](done <-chan K, routines ...<-chan T) <-chan T {
	wg := sync.WaitGroup{}

	finedInStream := make(chan T)

	convert := func(receivedChan <-chan T) {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case finedInStream <- <-receivedChan:
			}
		}
	}

	for _, routine := range routines {
		wg.Add(1)
		go convert(routine)
	}

	go func() {
		wg.Wait()
		defer close(finedInStream)
	}()

	return finedInStream
}

func fanOut[T Number, K bool](done <-chan K, stream <-chan T) []<-chan T {
	availableCpu := runtime.NumCPU()
	routines := make([]<-chan T, availableCpu)
	color.Cyan("Available CPU: %d\n", availableCpu)

	for i := range availableCpu {
		routines[i] = primeFinderAnalog(done, stream)
	}

	return routines
}

func primeFinderAnalog[K bool, T Number](done <-chan K, inputChan <-chan T) <-chan T {

	findPrime := func(value T) bool {

		for i := value / 2; i > 1; i++ {
			if value%i == 0 {
				return true
			}
		}

		return false
	}

	primeChan := make(chan T)

	go func() {
		defer close(primeChan)

		for {
			value := <-orDone(done, inputChan)

			if findPrime(value) {
				primeChan <- value
			}
		}

	}()

	//go func() {
	//	defer close(primeChan)
	//
	//	for {
	//		select {
	//		case <-done:
	//			return
	//		case pred := <-inputChan:
	//			if findPrime(pred) {
	//				primeChan <- pred
	//			}
	//		}
	//	}
	//}()

	return primeChan
}

func repeatAnalog[K any, T Number](done <-chan K, f func() T) <-chan T {
	output := make(chan T)

	go func() {
		defer close(output)
		for {
			select {
			case <-done:
				return
			case output <- f():
			}
		}
	}()

	return output
}

func takeAnalog[K any, T Number](done <-chan K, inputChan <-chan T, border int) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for i := 0; i < border; i++ {
			stream <- <-orDone(done, inputChan)
		}
	}()

	//go func() {
	//	defer close(stream)
	//
	//	for i := 0; i < border; i++ {
	//		select {
	//		case <-done:
	//			return
	//		case stream <- <-inputChan:
	//		}
	//	}
	//
	//}()

	return stream
}

func orDone[K any, T Number](done <-chan K, receivedChan <-chan T) <-chan T {

	resultChan := make(chan T)

	go func() {
		defer close(resultChan)

		for {
			select {
			case <-done:
				return
			case v, ok := <-receivedChan:
				if !ok {
					return
				}
				select {
				case resultChan <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return resultChan

}

func repeatFunction[T any, K any](done <-chan K, f func() T) <-chan T {
	outputChan := make(chan T)

	go func() {
		defer close(outputChan)

		for {
			select {
			case <-done:
				return
			case outputChan <- f():
			}
		}
	}()

	return outputChan
}

func repeatFunc[T Number, K any](done <-chan K, f func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- f():
			}
		}
	}()

	return stream
}

func take[T Number, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func primeFinder[T Number, K any](done <-chan K, stream <-chan T) <-chan T {

	isPrime := func(numb T) bool {
		for i := numb - 1; i > 1; i-- {
			if numb%i == 0 {
				return false
			}
		}
		return true
	}

	primed := make(chan T)

	go func() {
		defer close(primed)
		for {
			select {
			case <-done:
				return
			case iterableValue := <-stream:
				if isPrime(iterableValue) {
					primed <- iterableValue
				}
			}
		}
	}()

	return primed
}
