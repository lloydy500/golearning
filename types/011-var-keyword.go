package main

import (
	"fmt"
)

var a int // declare a variable with identifier 'a'
// and that the VARIABLE with the INDENTIFIER 'a' is of
// TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to 'a'
// false for booleans , 0.0 for floars, "" for strings and nil for pointers
// functions interfaces, slices, channels and maps.

// DECLARE and ASSIGN == INITIALIZATIONs

// go is a STATIC programming language, not DYNAMIC
// variable is DECLARED to hold a VALUE of a certain TYPE

var z = 55 // use var for PACKAGE scope declaration

func main() {

	// DECLARE the variable then ASSIGN value
	// short declaration operator

	x := 42

	fmt.Println(x)
	var y = 43 // this is bad practice, limit the scope as much as possible
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", y)
}
