package main

import "fmt"

// create a slice of a slice of strin ([][]string). Store the following data in t
// the multi-dimensional slice:
// "James", "Bond", "Shaken not stirred."
// "Miss", "Moneypenny", "helloooo James."
// Range over the records, then range over the data in each record.
func main() {
	xs1 := []string{"James", "Bond", "Shaken not stirred."}
	xs2 := []string{"Miss", "Moneypenny", "helloooo James."}
	xxs := [][]string{xs1, xs2}

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, val)
		}
	}
}
