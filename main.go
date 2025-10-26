package main

import "fmt"

func main() {

	bilik := newBill1("Second one")

	bilik.addItem("minon", 3.6)
	bilik.updateTip(10)

	res := bilik.format()

	fmt.Print(res)

	//myBill := newBill("First bill")
	//
	//fmt.Println(myBill)
	//
	//total := myBill.format()
	//fmt.Println(total)
}
