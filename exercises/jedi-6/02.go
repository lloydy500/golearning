package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 1, 45, 1}
	a := foo(xi...)
	b := bar(xi)

	fmt.Println(a, b)
}

// create a func with the identifier foo that@
// takes in anvariadic parameter of type int
// pass in a value of type []int into your func
// returns the sum of all values of type in passed in

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all the values of type int passed in

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// print the results
