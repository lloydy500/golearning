package main

import "fmt"

// a struct is a COMPOSITE DATA TYPE aka. COMPLEX DATA TYPE aka. AGGREGATE DATA TYPE
// make a new struct called "person"
type person struct {
	first string
	last  string
}

func main() {
	// we can create new VALUES of TYPE person!
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Printf("%T\n", p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
