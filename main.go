package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
)

type account struct {
	name     string
	password string
	url      string
}

func (a *account) generatePassword(length int) string {
	initialRunes := []rune(strings.ReplaceAll("abc d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1234567890-*@?!_", " ", ""))
	result := ""

	for i := 0; i < length; i++ {
		result += string(initialRunes[rand.Intn(len(initialRunes))])
	}
	a.password = result

	return result
}

func newUser(name, password, urlString string) (*account, error) {

	if len(name) == 0 {
		return nil, errors.New("THE_NAME_IS_EMPTY")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("url does not contain a valid account")
	}

	if len(password) == 0 {
		user := account{name: name, url: urlString}
		user.generatePassword(16)
		return &user, nil
	}

	return &account{name: name, password: password, url: urlString}, nil
}

func main() {

	fmt.Println(newUser("Roman", "", "https://roman.ua"))

	//newUser := account{}
	//
	//newUser.generatePassword(16)
	//
	//fmt.Println(newUser.password)

	//fmt.Println(strings.Repeat("-", 20))
	//fmt.Println(generatePassword(10))
}

func generatePassword(n int) string {

	initialRunes := []rune(strings.ReplaceAll("abc d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1234567890-*@?!_", " ", ""))
	result := ""

	for i := 0; i < n; i++ {
		result += string(initialRunes[rand.Intn(len(initialRunes))])
	}

	return result
}
