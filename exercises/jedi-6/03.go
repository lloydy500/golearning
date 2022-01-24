package main

import "fmt"

// use the defer keyword to show that the
// deferred func runs after the func containing it exists.

func main() {
	defer close()
	open()
}

func open() {
	fmt.Println("Opening file...")
}
func close() {
	fmt.Println("Closing file...")
}
