package main

import "fmt"

// In Go, a pointer is a variable that holds the memory address of another variable.
// POINTER: It allows you to indirectly access and modify the value of the variable it points to. 
// By using pointers, you can avoid copying large data structures, making your code more efficient and concise.

func updateName(x *string) {
	// fmt.Println("The variable name has been updated to", x)
	// return x
	*x = "wedge"
}

func main() {
	name := "Opaleti Oluwatobi"
	var p *string = &name
	m := &name
	fmt.Println(&m)
	fmt.Println(*p) 
	updateName(p)
	fmt.Println("Original variable name", name)
	fmt.Println(*p)
}