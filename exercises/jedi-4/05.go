package main

import "fmt"

// starting with below slice
// use APPEND AND SLICING  to get the values here which you should then print:
// [42,43,44,48,49,50,51]
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
