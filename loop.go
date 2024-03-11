package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	// Alternative method to the above code
	for i := 0; i<5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	// Looping through an array of strings
	names := []string{"John", "Doe", "Jane", "Doe"}
	for x := 0; x < len(names); x++ {
		fmt.Println(names[x])
	}

	// Looping through a string array and formatting the output using "Printf"
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
}