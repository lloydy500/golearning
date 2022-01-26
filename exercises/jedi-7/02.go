package main

import "fmt"

func main() {
	p1 := person{
		"Ethlebert",
		"Cumberdale",
		42,
	}
	fmt.Println(p1)
	changeMe(&p1) // this is get the address
	fmt.Println(p1)
}

// create person struct
type person struct {
	first string
	last  string
	age   int
}

// create a func called "changeMe" with a *person as a param
// change the value stored at the *person address

func changeMe(p *person) { // this takes a pointer to a person
	(*p).age *= 2
	// or just p.age compiler will convert
}
