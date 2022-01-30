package main

import "fmt"

// access every element in the slice using a for loop
func main() {
	x := []int{4, 5, 7, 8, 42}
	x = append(x, 77, 88, 99, 114, 1014)
	fmt.Println(x)

	y := []int{234, 456, 987}
	x = append(x, y...)
	fmt.Println(x)
	// T... is 'unfurl values' from a given data structure and is diffent to ...T in the docs is a variadic param
	// this is a param that takes unlimited number of args
}
