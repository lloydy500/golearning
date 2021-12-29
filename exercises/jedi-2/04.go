package main

import "fmt"

func main() {
	// Write a program that:
	// assigns an int to a VARIABLE
	a := 21
	// prints that int in DECIMAL, BINARY and HEX
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	// shifts the bits of that int over 1 position to the left
	// and assigns that to a variable
	b := a << 1
	// prints that variable in DECIMAL, BINARY, and HEX.
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)

}
