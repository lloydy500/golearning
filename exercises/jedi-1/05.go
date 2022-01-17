package main

import "fmt"

type bigballs int // int is the UNDERLYING type
// at the package level SCOPE, create a VARIABLE with the INDENTIFIER "y"
// the variable TYPE should be of the type of your UNDERLYING TYPE OF X

var x bigballs
var y int

// Use CONVERSION to convert the TYPE of the VALUE stored in x
// to the UNDERLYING TYPE
// ASSIGN that value to y
// print out the value stored in "y"
// print our the TYPE of "y"
func main() {
	x = 42
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
