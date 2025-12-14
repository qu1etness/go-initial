package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

type account struct {
	name     string
	password string
	url      string
}

// Вбудовування (Унаслідування)

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

type AccountWithTimeStampAnalog struct {
	createdAt time.Time
	updatedAt time.Time
	acc       account
}

func (a *account) outputPassword() {
	fmt.Println(a.password)
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

	user := account{name: name, url: urlString, password: password}

	if len(password) == 0 {
		user.generatePassword(16)
	}

	return &user, nil
}

func NewUserWithTimeStamp(name, password, urlString string) (*AccountWithTimeStamp, error) {

	if name == "" {
		return nil, errors.New("THE_NAME_IS_EMPTY")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("THE_URL_IS_INVALID")
	}

	user := AccountWithTimeStamp{
		account:   account{name: name, url: urlString, password: password},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if password == "" {
		user.generatePassword(16)
	}

	return &user, nil
}

func generatePassword(n int) string {

	initialRunes := []rune(strings.ReplaceAll("abc d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1234567890-*@?!_", " ", ""))
	result := ""

	for i := 0; i < n; i++ {
		result += string(initialRunes[rand.Intn(len(initialRunes))])
	}

	return result
}
