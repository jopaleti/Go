package main

import (
	"fmt"
	"math"
	"strconv"
)

// Struct
// Defining person struct
type Person struct {
	// name string
	// age int
	// gender string

	name, gender string
	age int
}

func adder() func(int) int {
	sum := 0
	return func (x int) int {
		sum += x
		return sum
	}
}

// Greeting method
func (p Person) greet() string {
	return "Hello " + p.name + " your age is " + strconv.Itoa(p.age) + " and you are a " + p.gender
}

// Pointer receiver
func (p *Person) hasbirthday() {
	p.age++
}

// function changeName(pointer receiver)

func(p *Person) changeName(name string) {
	if p.gender == "female" {
		return
	} else {
		p.name = name
	}
}

// Define interface
type Shape interface {
	area() float64
}

//  Define different struct
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

// Area method
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Defining function that return area
func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	// INITIALIZING CIRCLE AND RECTANGLE FUNCTION
	circle := Circle{x:0, y:0, radius:5}
	rectangle := Rectangle{width: 20, height: 40}
	
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

	sum := adder()
	for i := 0; i<10; i++ {
		fmt.Println(sum(i))
	}

	// Understanding Pointer
	x := 5
	fmt.Println(&x)

	// Initialising the person struct
	person := Person{name: "Oluwatobi", age: 21, gender: "male"}
	person.hasbirthday()
	person.changeName("Johnson")
	fmt.Println(person.name)
	fmt.Println(person.greet())

	// fmt.Println("Hello World!");
	// fmt.Println("2+2 = ",1+1)
	// fmt.Println(true || false)

	var a = "initial"
	var b, c int = 1, 3
	fmt.Println(a, b, c)
	const n = 5000

	f := "apple"
	fmt.Println(f)
	fmt.Println(math.Sin(60))
	fmt.Println(n)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for n := 0; n <= 6; n++ {
		if n%2 == 0 {
			continue
		} else {
			fmt.Println(n)
		}
		// fmt.Println(n)
	}
}