package main

import "fmt"

// create a slice to store the names of all the states in the US.
// Use MAKE and APPEND to do this. Goal is to not create underlying ARRAY
// more than once. what is length of slice? what is capacity?
// printout all the values and their indexes without using the range clause

func main() {
	x := make([]string, 0, 50)
	x = append(x, "Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana",
		"Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming")
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i <= 49; i++ {
		fmt.Println(x[i])
	}
}
