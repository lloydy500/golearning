package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string // use caps so we can access these fields
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{
		p1, p2,
	}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
