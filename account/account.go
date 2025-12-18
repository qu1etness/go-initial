package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Account) outputPassword() {
	color.Blue(a.Password)
}

func (a *Account) generatePassword(length int) string {
	initialRunes := []rune(strings.ReplaceAll("abc d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1234567890-*@?!_", " ", ""))
	result := ""

	for i := 0; i < length; i++ {
		result += string(initialRunes[rand.Intn(len(initialRunes))])
	}
	a.Password = result

	return result
}

func NewUser(name, password, urlString string) (*Account, error) {

	if len(name) == 0 {
		return nil, errors.New("THE_NAME_IS_EMPTY")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("URL_IS_NOT_VALID")
	}

	user := &Account{Name: name, Url: urlString, Password: password}

	if len(password) == 0 {
		user.generatePassword(16)
	}

	value, isExist := reflect.TypeOf(user).Elem().FieldByName("name")
	fmt.Println("Reflect output: ", value.Tag, isExist)

	return user, nil
}

func (a *Account) ToBytes() []byte {
	byteSlice, _ := json.Marshal(a)
	return byteSlice
}

func generatePassword(n int) string {
	initialRunes := []rune(strings.ReplaceAll("abc d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1234567890-*@?!_", " ", ""))
	result := ""

	for i := 0; i < n; i++ {
		result += string(initialRunes[rand.Intn(len(initialRunes))])
	}

	return result
}
