package main

import "fmt"

// Pipelines

func main() {

	initialValues := []int{2, 4, 6, 7, 1}

	//	Stage 1
	dataChannel := sliceToChan2(initialValues)
	//	Stage 2
	finalChannel := sq2(dataChannel)
	//  Stage 3

	for value := range finalChannel {
		fmt.Println(value)
	}

}

func sq2(inputChan <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for value := range inputChan {
			out <- value * value
		}
	}()

	return out
}

func sliceToChan2(initialValues []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range initialValues {
			out <- v
		}
	}()

	return out
}

func sq(channel <-chan int) chan int {
	outputChannel := make(chan int)

	go func() {
		for value := range channel {
			outputChannel <- value * value
		}
		close(outputChannel)
	}()
	return outputChannel
}

func sliceToChannel(nums []int) <-chan int {
	outputChannel := make(chan int)

	go func() {
		for _, number := range nums {
			outputChannel <- number
		}
		close(outputChannel)
	}()

	return outputChannel
}
