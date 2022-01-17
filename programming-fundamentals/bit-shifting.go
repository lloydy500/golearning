package main

import (
	"fmt"
)

const (
	_ = iota // throw away the 0 iota. we don't need it here
	// kb = 1024
	// shift 1 over 10 to the left
	kb = 1 << (iota * 10) // iota == 1
	// shift 1 over 30 to the left
	mb = 1 << (iota * 10) // iota == 2
	// shift 1 over 40 to the left
	gb = 1 << (iota * 10) // iota == 3
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	// each time we shift over 10 times
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
