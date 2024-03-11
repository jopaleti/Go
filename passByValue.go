package main

import "fmt"

/**
Go is a pass by value language
Which means when a variable is passed into a function Go will make a copy of it
and store it into a new a block or space of memory, then it stores the pointer as well 
inside a new block of space or memory.
POINTER: It points to a 
*/

func updateName(x string) {
	x = "wedge"
}

func main() {
	name := "Oluwtobi"

	// Storing the memory address of variable name in variable m below

	m := &name

	updateName(name)

	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(*m)
}