package main

import (
	"fmt"
)

func main() {

	//fmt.Println(newUser("Roman", "", "https://roman.ua"))

	//newUser := account{}
	//
	//newUser.generatePassword(16)
	//
	//fmt.Println(newUser.password)

	//fmt.Println(strings.Repeat("-", 20))
	//fmt.Println(generatePassword(10))

	modernUser, err := NewUserWithTimeStamp("Ivan", "", "https://ivan.ua")

	if err != nil {
		panic(err)
	}

	fmt.Println(modernUser)

}
