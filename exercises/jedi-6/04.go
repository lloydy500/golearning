package main

import "fmt"

// create a user defined struct
// with the fields:
// first
// last
// age
// attach a method to type person with the identifier speak
// the method should have the person say their name and age
// create a vallue of type person
// call hte method from the value of type person

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I'm %v years old.\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "John",
		last:  "Smith",
		age:   29,
	}
	p1.speak()
}
