package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	resultArr := make([]float32, 0, 10)

	for {
		fmt.Println(" - Print \"q\" to exit")

		element, err := getInput("Add another element of array", reader)
		if err != nil {
			fmt.Println("Something went wrong!!!")
		}
		if element == "q" {
			break
		}
		parsedElement, err := strconv.ParseFloat(element, 32)

		if err != nil {
			fmt.Println("Something went wrong!!!")
		}

		resultArr = append(resultArr, float32(parsedElement))
	}

	var sum float32

	for _, value := range resultArr {
		sum += value
	}

	fmt.Printf("Array is  %v \n", resultArr)
	fmt.Printf("The sum of array is %5v \n", sum)
}

func getInput(placeholder string, r *bufio.Reader) (string, error) {
	fmt.Print(placeholder + ":")
	input, err := r.ReadString('\n')

	if err != nil {
		return "", err
	}

	trimmedStr := strings.TrimSpace(input)

	return trimmedStr, nil
}
