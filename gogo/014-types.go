package main

import (
	"fmt"
)

var a int

// we can create our own TYPE in go
type hotdog int

var b hotdog

func main() {

	a = 42
	b = 43
	//  can't do a = b. var a is type int, var b is type hotdog.
	a = int(b) // this is called CONVERSION. Other langauges call it CASTING
	// we have CONVERTED the b of TYPE hotdog to TYPE int and ASSIGNED to a
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
