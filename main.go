package main

import "fmt"

// For-select statement in Go allows a goroutine to wait on multiple communication operations.

func main() {
	// unbuffered channel charChannel := make(chan string)

	//	buffered channel
	charChannel := make(chan string, 3)
	chat := [3]string{"a", "b", "c"}

	for _, v := range chat {
		select {
		case charChannel <- v:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}
