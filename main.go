package main

import (
	"fmt"
)

func main() {
	infiniteCalculation()
}

func infiniteCalculation() {
	var enterValue uint8
	for {
		fmt.Println("---------Calculator--------")
		fmt.Println("Enter the value:")
		fmt.Println("1. Calculate new data")
		fmt.Println("2. Exit")
		fmt.Print("Enter value: ")
		_, err := fmt.Scan(&enterValue)

		if err != nil {
			fmt.Println("Error:", err)
			var discard string
			fmt.Scan(&discard)
			fmt.Println("Invalid input, try again")
			continue
		}

		switch enterValue {
		case 1:
			userMass, userHeight := getInputValue()
			index := calculateIndex(userMass, userHeight)
			showInfo(userMass, userHeight, index)
		case 2:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Please enter a valid option (1 or 2).")
		}
	}
}

func getInputValue() (userMass float32, userHeight float32) {
	fmt.Println("============Calculation============")
	fmt.Println("Enter a mass (in kg):")
	fmt.Scan(&userMass)
	fmt.Println("Enter a height (in metres):")
	fmt.Scan(&userHeight)
	return
}
func calculateIndex(userMass float32, userHeight float32) float32 {
	var index float32 = userMass / (userHeight * userHeight)
	return index
}
func showInfo(userMass float32, userHeight float32, index float32) {
	fmt.Printf("User with %.0f kg mass and %0.2f m height has %0.1f index \n", userMass, userHeight, index)
}

//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
