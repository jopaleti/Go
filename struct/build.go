package main

import (
	"fmt"
)

type bill struct {
	name string
	items map[string]float64
	tip float64 
}

// func  (b bill) struct_construct(name string) {
// 	b.name = name
// 	b.items = map[string]float64
// 	tip = 0
// }

// MAKING NEW BILL
func newBill(name string) bill {
	b := bill {
		name: name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip: 0,
	}
	return b
}

// FORMAT THE BILL
// Receiver function
func (s bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	for k, v := range s.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total)

	return fs
}

// Updating the tip: Pass the struct as a POINTER
func (b *bill) updateTip(tip float64) {
	// Since we are updating struct, it is not necessary to de reference it but we may de reference it
	// If we do not de refence it, go will automatically de reference it for us 
	b.tip = tip
}

// Add an item to the bill
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}