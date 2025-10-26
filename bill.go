package main

import "fmt"

type bill struct {
	name  string
	items map[string]float32
	tip   float32
}

func newBill(name string) bill {
	exemplar := bill{
		name:  name,
		items: map[string]float32{"romik": 7.5, "olko": 7.4, "kris": 5.6},
		tip:   0,
	}
	return exemplar
}

func (b *bill) format() string {
	result := "The value is: \n"
	var total float32

	for k, v := range b.items {
		result += fmt.Sprintf("%-25v ...%v \n", k+":", v)
		total += v
	}

	result += fmt.Sprintf("%-25v ...%0.2f", "total:", total)
	return result
}
