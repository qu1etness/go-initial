package main

import (
	"fmt"
	"strings"
)

func main() {

	initialName, initialSurname := getInitials("Roman Kostiv")
	fmt.Println(initialName, initialSurname)

}

func getInitials(str string) (string, string) {

	splited := strings.Split(str, " ")

	if len(splited) < 2 {
		return string(splited[0][0]), "_"
	}
	return string(splited[0][0]), string(splited[1][0])
}
