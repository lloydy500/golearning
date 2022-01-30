package main

import "fmt"

// starting with below slice
// APPEND to that SLICE this value: 52
// print out the slice
// in ONE STATEMENT append to that slice the values: 53,54,55
// print out the the slice
// append to the slice this slice
// y:= []int{56,57,58,59,60}
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}

//  ezz
