package main

import "fmt"

func main() {
	firstChanel := make(chan string)

	go func() {
		firstChanel <- "Concurrency in Go!"
	}()

	msg := <-firstChanel

	fmt.Println(msg)
}
