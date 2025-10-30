package main

import (
	"fmt"
)

func fibonacci() func() int {
	initial := 1

	return func() int {
		initial += initial
		return initial
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
