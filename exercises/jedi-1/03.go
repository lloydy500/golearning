package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

// use fmt.Sprintf to print all of the VALUES to one single string.
// ASSIGN the returned value of TYPE string using the short declaration operator to a
// VARIABLE with the IDENTIFIER "s"
func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
