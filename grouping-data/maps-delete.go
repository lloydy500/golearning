package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	delete(m, "James")

	fmt.Println(m)
	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian Flemming"])
	if v, ok := m["Miss Moneypeny"]; ok {
		fmt.Println("value", v)
		delete(m, "Miss Moneypenny")
	}
}
