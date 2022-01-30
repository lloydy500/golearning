package main

import "fmt"

// a struct is a COMPOSITE DATA TYPE aka. COMPLEX DATA TYPE aka. AGGREGATE DATA TYPE
// make a new struct called "person"
// here person is the INNER TYPE
type person struct {
	first string
	last  string
}

// we can take person and embed it into secretAgent
// add a license to kilL
type secretAgent struct {
	person
	ltk bool
}

func main() {
	// we can create new VALUES of TYPE person!
	// the INNER TYPE gets PROMOTED to the OUTER TYPE
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Printf("%T\n", sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	fmt.Println(p2.first, p2.last)
}
