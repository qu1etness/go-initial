package main

import "fmt"

// Pipelines

func main() {

	initialValues := []int{2, 4, 6, 7, 1}

	//	Stage 1
	dataChannel := sliceToChannel(initialValues)
	//	Stage 2
	finalChannel := sq(dataChannel)
	//  Stage 3

	for value := range finalChannel {
		fmt.Println(value)
	}

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
