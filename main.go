package main

import "fmt"

func main() {
	firstChanel := make(chan string)
	secondChanel := make(chan string)

	go func() {
		firstChanel <- "Concurrency in Go!"
	}()

	go func() {
		secondChanel <- "Is fun!"
	}()

	select {
	case firstChanelMsg := <-firstChanel:
		fmt.Println(firstChanelMsg)
	case secondChanelMsg := <-secondChanel:
		fmt.Println(secondChanelMsg)
	}
}
