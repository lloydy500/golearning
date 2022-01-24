package main

import "fmt"

// Print every rune code point of the
// uppercase alphabet three times.

// lookup ASCII and see uppercase alphabet is 65-90
func main() {
	count := 1
	for i := 65; i < 91; i++ {
		fmt.Printf("%v\n", count)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
		count++
	}
}
