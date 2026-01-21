package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	input := []int{1, 2, 3, 4, 5, 6}
	output := make([]int, len(input))
	wg := sync.WaitGroup{}

	for i, v := range input {
		wg.Add(1)
		go processData(&wg, &output[i], v)
	}

	wg.Wait()

	fmt.Println(input)
	fmt.Println(output)

	fmt.Println(time.Since(start))

}

func processData(wg *sync.WaitGroup, result *int, initial int) {
	defer wg.Done()

	outputValue := process(initial)

	*result = outputValue
}

func process(initial int) int {
	time.Sleep(2 * time.Second)
	return initial * 2
}
