package main

import "fmt"

// create TYPED and UNTYPED CONSTANTS.
// Print the values of the contants.

const (
	a     = "lloyd" // UNTYPED CONSTANT
	b int = 29      // TYPED CONSTANT
)

func main() {
	fmt.Println(a, "is", b)
}
