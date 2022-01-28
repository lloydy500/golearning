package main

// get this code running

import (
	"fmt"
)

func main() {
	// cs := make(<-chan int) // channel need to be bidirectional
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
