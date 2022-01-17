package main

import "fmt"

// use var to DECLARE three variables. The variables should have
// package level scope. Use the following INDENTIFIERS for the
// variables and make sure the variables are of the following TYPE
// (meaning they can store VALUES of that TYPE)
// x int, y string, z bool

var x int
var y string
var z bool

// in func main print out the values for each identifier
// the compiler assigned values to the variables. What are these values called?
func main() {
	fmt.Println(x, y, z) // <-- when the compiler assigns values
	// to these variables, we call these values "Zero values"

}
