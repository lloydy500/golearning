package main

import (
	"fmt"
)

// create your own type "person" which will have an
// underlying type of struct so that it can store
// the following data:
// first_name
// last_name
// favorite ice cream flavors

// create two VALUES of TYPE person. Print out the values,
// ranging over the elements in the slice

type person struct {
	first_name     string
	last_name      string
	fav_ice_creams []string
}

func main() {
	p1 := person{
		first_name:     "James",
		last_name:      "Smith",
		fav_ice_creams: []string{"chocolate", "praline", "raspberry"},
	}
	p2 := person{
		first_name:     "Ju",
		last_name:      "Ju",
		fav_ice_creams: []string{"banana", "mango", "papaya"},
	}

	fmt.Printf("The first person is %v %v and their favourite ice creams are:\n", p1.first_name, p1.last_name)

	for _, val := range p1.fav_ice_creams {
		fmt.Printf("%v\n", val)
	}
	fmt.Printf("The second person is %v %v and their favourite ice creams are:\n", p2.first_name, p2.last_name)

	for _, val := range p2.fav_ice_creams {
		fmt.Printf("%v\n", val)
	}
}
