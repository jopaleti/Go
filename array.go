package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	var ages [3]int = [3]int{20, 25, 30}

	// Shorthand declaration of variable
	names := [4]string{"John", "Doe", "Jane", "Doe"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices
	scores := []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// Slice ranges
	fmt.Println(scores[2:])

	// Strings module
	// Note strings.Contains is case sensitive
	greeting := "Hello world I am a programmer"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(greeting)
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, ""))

	// Sorting Package
	Ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(Ages)
	fmt.Println(Ages)

	Index := sort.SearchInts(Ages, 50)
	fmt.Println(Index)

	Names := []string{"John", "Doe", "Jane", "Doe"}
	sort.Strings(Names)
	fmt.Println(Names)
}