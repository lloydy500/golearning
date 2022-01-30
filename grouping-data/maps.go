package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	} // look at the COMPOSITE LITERAL - the TYPE is "map[string]int"
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabus"])

	v, ok := m["Barnabus"]
	fmt.Println(v)
	fmt.Println(ok)

	// this is the COMMA, OK IDIOM
	// very common in go
	if v, ok := m["Barnabus"]; ok {
		fmt.Println("This is the if print", v)
	}

	if v, ok := m["James"]; ok {
		fmt.Println("This is the second if print", v)
	}

}
