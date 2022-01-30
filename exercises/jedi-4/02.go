package main

import "fmt"

func main() {
	// using a COMPOSITE LITERAL
	// create a SLICE of TYPE int
	// assign 10 values
	// range over the slice and print the values out.
	// using format printing print out the TYPE of the SLICE

	x := []int{22, 12, 42, 55, 983, 91827, 21, 325, 342, 23}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}
