package main

import "fmt"

//  a "person" was just an identifier of a struct with
//  the specific fields first, last, age
type person struct {
	first string
	last  string
	age   int
}

func main() {
	// this is an anonymous struct
	// we can do this where we want to keep code lean and readable
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Harry",
		last:  "Potter",
		age:   13,
	}
	fmt.Println(p1)
}
