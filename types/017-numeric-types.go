package main

import "fmt"

var x int
var y float64

var z int8

// integers have no decimals
// floating points do
// floating points are called REAL NUMBERS
func main() {
	x = 42 // putting a decimal here will break the code
	// because we have declared above that x is of TYPE int
	y = 42.34534
	z = -128 // all the way up to positive 127
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

}

// BYTE is an alias for uint-8 (unsigned integer-8)
// RUNE (or character) is an alias for int32. for each character
// in utf-8 each character need up to 4 BYTES (OR 32 BITS)
