package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }
// we've attached the speak function to any VALUE of type r
func (s secretAgent) speak() {
	fmt.Println("The name's", s.last, ".", s.first, s.last, ".")
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}
