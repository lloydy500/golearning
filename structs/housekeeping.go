package main

import "fmt"

// look at the similarities below
type foo int

var x int
var y foo

const bar int = 42

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}
	fmt.Println(p1)
	y = 42
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
}

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS
