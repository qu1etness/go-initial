package main

import (
	"fmt"
	"time"
)

// Done-select statement in Go allows a goroutine to wait on multiple communication operations.

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Work in progress...")
		}
	}
}

func main() {

	startTime := time.Now()
	doneChan := make(chan bool)

	go doWork(doneChan)

	time.Sleep(2 * time.Second)

	close(doneChan)

	duration := time.Since(startTime)
	fmt.Println(duration)
}
