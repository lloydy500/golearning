package main

import "fmt"

// a func is a TYPE like any other, so it
// can be assigned to a variable
func main() {
	f := func() {
		fmt.Println("my first func expression")
	}
	f()
	o := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	o(1984)
}
