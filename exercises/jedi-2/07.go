package main

import "fmt"

// print all the numbers from 33 to 122 as numbers
// print all the numbers from 33 to 122 as ASCII letters

func main() {
	x := 33

	for x <= 122 {
		fmt.Printf("%v\t%#x\t%#U\n", x, x, x)
		x++
	}
}
