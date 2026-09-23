package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		ch <- 100
	}()
	fmt.Println(<-ch)
}
