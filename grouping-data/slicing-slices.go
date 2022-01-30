package main

import "fmt"

// access every element in the slice using a for loop
func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1:3]) // same as python

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
	// OR
	for i, v := range x {
		fmt.Println(i, v)
	}
}
