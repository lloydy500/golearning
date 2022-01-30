package main

import "fmt"

func main() {
	// create a map with a key of TYPE string whichh is a
	// persons "last_first" name, and a value of TYPE
	// []string which stores their favorite things
	// store 3 records and print out all values and indices in the slice

	x := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "martinis", "women"},
		"miss_moneypenny": []string{"James Bond", "Literature", "computer science"},
		"no_dr":           []string{"being evil", "ice cream", "sunsets"},
	}

	x["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for k, v := range x {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\t %v.\t %v\n", i, val)
		}
	}
}
