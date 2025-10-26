package main

import "fmt"

type bill1 struct {
	name  string
	items map[string]float32
	tip   float32
}

func newBill1(name string) bill1 {
	example := bill1{
		name:  name,
		items: map[string]float32{"romik": 7.5, "olko": 7.4, "kris": 5.6},
		tip:   0,
	}
	return example
}

func (b *bill1) format() string {
	result := "--------Info about bill1 structure--------\n"
	var total float32

	for key, value := range b.items {
		result += fmt.Sprintf("%-25v is %.2f f \n", key+":", value)
		total += value
	}

	result += fmt.Sprintf("%-25v is %0.2f", "total:", total)

	return result
}

func (b *bill1) updateTip(tip float32) float32 {
	b.tip = tip // -> (*b).tip
	return b.tip
}

func (b *bill1) addItem(key string, value float32) float32 {
	b.items[key] = value
	return b.items[key]
}
