package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type bookmarkMap = map[string]string

func main() {
	manageBookmarks()
}

func getInput(placeholder string, r *bufio.Reader) (string, error) {
	fmt.Print(placeholder + ":")

	input, err := r.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func manageBookmarks() {
	bookmark := make(bookmarkMap, 5)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=========== Manage Bookmarks ============")
Menu:
	for {
		fmt.Println("1. Look at bookmarks")
		fmt.Println("2. Create new bookmark")
		fmt.Println("3. Delete bookmark by key")
		fmt.Println("4. Exit")

		option, err := getInput("Enter ur option", reader)

		if err != nil {
			fmt.Println(errors.New("something went wrong with input value"))
		}

		switch option {
		case "1":
			lookThoughtBookmarks(bookmark)
		case "2":
			key, keyError := getInput("Enter bookmark's key", reader)
			value, valueError := getInput("Enter bookmark's value", reader)

			if valueError != nil || keyError != nil {
				fmt.Println(valueError, keyError)
				fmt.Println("Invalid key or value input. Try again...")
				continue
			}
			fmt.Println(addNewBookmark(bookmark, key, value))
		case "3":
			deleteKey, err := getInput("Enter the key, you gonna delete", reader)
			if err != nil {
				fmt.Println("Invalid key input. Try again...")
			}
			fmt.Println(deleteBookmarkByKey(bookmark, deleteKey))
		case "4":
			fmt.Println("Goodbye👋")
			break Menu
		default:
			fmt.Println("Invalid option. Please enter the value from the list")
		}
	}

}

func deleteBookmarkByKey(b bookmarkMap, key string) string {
	_, isExist := b[key]
	if isExist {
		delete(b, key)
		return fmt.Sprintf("bookmark with %v key was successfuly deleted", key)
	}
	return "There is no bookmark with this key"
}

func addNewBookmark(b bookmarkMap, key string, value string) string {
	b[key] = value
	return fmt.Sprintf("New bookmark added with key %v and value %v\n", key, value)
}

func lookThoughtBookmarks(b bookmarkMap) {
	fmt.Print("------ ALl the bookmarks ------ \n\n")
	for key, value := range b {
		fmt.Printf("In %v bookmark stores: %5v\n", key, value)
	}
}
