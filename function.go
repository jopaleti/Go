package main

import (
	"fmt"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

// Call stack explanation in programming
// it follows the LIFO (Last In First Out) principle and similar to a stack data structure
func f1(){
	fmt.Println("This is f1")
}
func f2(){
	fmt.Println("This is f2")
}

func f3(){
	fmt.Println("This is f3")
	f1()
	f2()
}

func main() {
	names := []string{"John", "Doe", "Jane", "Doe"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)
}







