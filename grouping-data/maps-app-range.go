package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	} // look at the COMPOSITE LITERAL - the TYPE is "map[string]int"

	m["todd"] = 33 // add a new item

	for k, v := range m {
		fmt.Println(k, v)
	}
	xi := []int{4, 5, 6, 7, 8, 2, 23}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
