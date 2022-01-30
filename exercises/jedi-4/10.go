package main

import "fmt"

func main() {
	// delete a record
	x := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "martinis", "women"},
		"miss_moneypenny": []string{"James Bond", "Literature", "computer science"},
		"no_dr":           []string{"being evil", "ice cream", "sunsets"},
	}

	x["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	delete(x, "no_dr") //<----
	for k, v := range x {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\t %v.\t %v\n", i, val)
		}
	}
}
