package main

import "fmt"

// iota is a PREDECLARED IDENTIFIER and is a KEYWORD in the language
// specification
const (
	a = iota
	b
	c
)
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
