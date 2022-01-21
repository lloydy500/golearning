package main

import "fmt"

// take the code from previous exercise
// store the values of TYPE person in a map with
// the key of last_name. Access each value in the map.
// Print out the values, ranging over the slice.
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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	fmt.Println("-------")

	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, val := range v.fav_ice_creams {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
