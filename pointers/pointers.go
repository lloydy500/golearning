package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)  // value
	fmt.Println(&a) // & what is the address :)?

	fmt.Printf("%T\n", a) // int
	// different type
	fmt.Printf("%T\n", &a) // *int, pointer to an int

	// var b int = &a can't init an int and assign to a pointer

	// but we can share addresses where values are stored
	b := &a
	fmt.Println(b)
	// "*int" is a pointer to an int
	// *b is de-referencing an address, giving a value
	fmt.Println(*b)  // 42
	fmt.Println(*&a) // get the address and de-ref, 42
	*b = 43          // change the value at the address of b
	fmt.Println(a)   // same as address of a!!
}
