package main

import (
	"fmt"
	"time"
)

func routine(num string) {
	fmt.Println(num)
}

func main() {

	go routine("1")
	go routine("2")
	go routine("3")

	time.Sleep(time.Second * 2)

	fmt.Println("Main function")
}
