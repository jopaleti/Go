package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{"soup": 5.55,"stew": 4.78,"salad": 10.00}

	fmt.Println(menu)

	// Looping through the map(dictionary) and printing out the key - value pair
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
}