package main

import "fmt"

// Create your own type. have the underlying tyoe be an int.
type bigballs int // int is the UNDERLYING type

// Create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the "VAR"
// keyword.
var x bigballs

func main() {
	// in func main, print out value of variable x
	fmt.Printf("%v\n", x)
	// print type of variable x
	fmt.Printf("%T\n", x)
	// assign 42 to the VARIABLE x using the = OPERATOR
	x = 42
	// print out the value of the variable x
	fmt.Printf("%v\n", x)
}
