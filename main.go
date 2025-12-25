package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-initial/account"
	"go-initial/files"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	accountMenu(reader)
	//reader := bufio.NewReader(os.Stdin)
	//accountMenu(reader)
}

func createAccount(reader *bufio.Reader, recentVault *account.Vault) {
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

	recentVault.AddAccount(*user)
}

func findAccount(reader *bufio.Reader, recentVault *account.Vault) {

	researchingName, err := getInput("Enter account name", reader)
	if err != nil {
		color.Red("Failed to read from stdin:", err)
		return
	}

	recentVault.FindAccount(researchingName)
}

func deleteAccount(reader *bufio.Reader, recentVault *account.Vault) {

	deleteName, err := getInput("Enter account name", reader)

	if err != nil {
		color.Red("Failed to read from stdin:", err)
		return
	}

	recentVault.RemoveAccount(deleteName)
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

func accountMenu(reader *bufio.Reader) {

	titleStyle := color.New(color.FgBlack, color.BgCyan).PrintfFunc()
	errorStyle := color.New(color.FgRed, color.Bold).PrintfFunc()

	currentVault, _ := account.NewVault()

Main:
	for {
		titleStyle("============Account Menu============")
		fmt.Println("\n1. Create Account")
		fmt.Println("2. Find Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Show all accounts")
		fmt.Println("5. Exit")
		titleStyle("====================================")

		option, err := getInput("\nChoose ur option", reader)

		if err != nil {
			errorStyle("Failed to read from stdin:", err)
		}

		switch option {
		case "1":
			createAccount(reader, currentVault)
		case "2":
			findAccount(reader, currentVault)
		case "3":
			deleteAccount(reader, currentVault)
		case "4":
			currentVault.ShowAllAccounts()
		case "5":
			data, _ := json.Marshal(currentVault)

			files.WriteFile(data, "data.json")
			fmt.Println("Byee!")
			break Main
		default:
			errorStyle("Invalid option. Please try again.\n")
		}
	}
}
