package main

import (
	"fmt"
)

// using the short declaration operator, ASSIGN thesen VALUES
// to VARIABLES with the IDENTIFIERS X,Y AND Z.
// a. 42, b. "James Bond, c. true"

func main() {
	x := 42
	y := "James Bond"
	z := true

	// print using a single print statement
	fmt.Println(x, y, z)

	// print using multiple print statements
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
