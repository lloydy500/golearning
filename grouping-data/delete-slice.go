package main

import "fmt"

// access every element in the slice using a for loop
func main() {
	x := []int{4, 5, 7, 8, 42}
	x = append(x, 77, 88, 99, 114, 1014)
	fmt.Println(x)

	y := []int{234, 456, 987}
	x = append(x, y...)
	// delete 7 and 8 from the slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
