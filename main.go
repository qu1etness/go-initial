package main

import (
	"bufio"
	"fmt"
	"go-initial/account"
	"go-initial/files"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	createAccount()
}

func createAccount() {
	reader := bufio.NewReader(os.Stdin)

	name, err := getInput("Enter account name", reader)
	if err != nil {
		color.Red("Failed to read from stdin:", err)
		return
	}

	url, err := getInput("Enter account URL", reader)
	if err != nil {
		color.Red("Failed to read from stdin:", err)
		return
	}

	password, err := getInput("Enter account password (leave empty to generate)", reader)
	if err != nil {
		color.Red("Failed to read from stdin:", err)
		return
	}

	user, err := account.NewUser(name, password, url)
	if err != nil {
		color.Red("Failed to read from stdin:", err)
		return
	}

	files.WriteFile(user.ToBytes(), "data.json")
}

func getInput(label string, reader *bufio.Reader) (string, error) {

	fmt.Print(label + ": ")
	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", err
	}

	return strings.TrimSpace(value), nil
}
