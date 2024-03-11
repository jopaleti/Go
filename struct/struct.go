package main

import "fmt"

// STRUCT: It's a blueprint which describe the type of data

func main() {
	myBill := newBill("Opaleti Oluwatobi")
	myBill.updateTip(1112)
	myBill.addItem("Orange", 200)
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}