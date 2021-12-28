package main

import (
	"fmt"
)

var y = 911

func main() {

	// DECLARE a VARIABLE to be of a certain TYPE
	// and the ASSIGN a VALUE of that TYPE to the value
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%T\t%b\t%x", y, y, y)

	// print it all to a string and assign to s
	s := fmt.Sprintf("%T\t%b\t%x", y, y, y)
	// then print that s VARIABLE
	fmt.Println(s)
}
