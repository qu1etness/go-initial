package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(placeholder string, reader *bufio.Reader) (string, error) {
	fmt.Print(placeholder + ":")
	value, err := reader.ReadString('\n')
	return strings.TrimSpace(value), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name", reader)
	b := newBill(name)
	return b
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===Bill manager===")
	fmt.Println("Choose one option: ")
	fmt.Println("1. add item")
	fmt.Println("2. save bill")
	fmt.Println("3. add tip")
	fmt.Println("0. exit")
	option, _ := getInput("Enter a value", reader)

	switch option {
	case "1":
		fmt.Println(b)
	case "2":
		fmt.Println(b.items)
	case "3":
		fmt.Println(b.tip)
	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}
}

func main() {
	bill := createBill()
	promptOptions(&bill)

}
