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

type hotdog int

// func (r receiver) identifier(parameters) (return(s)) { code }
// we've attached the speak function to any VALUE of type r
func (s secretAgent) speak() {
	fmt.Println("The name's", s.last, ".", s.first, s.last, ". I'm a secret agent.")

}

func (p person) speak() {
	fmt.Println("The name's", p.first, p.last, ". I'm a secret person.")

}

//keyword identifier type
// a VALUE can be of more than one TYPE
// if it has the speak() method, then its human type
type human interface {
	speak()
}

// func bar takes an argument,h, of TYPE human and prints it
func bar(h human) {
	switch h.(type) {
	// assertion: this is how we ASSERT that h is of type person
	case person:
		fmt.Println("I was passed into barrrrr.", h.(person).first)
	case secretAgent:
	}
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
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}
	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	bar(sa1)
	bar(sa2)
	bar(p1)
	//conversion

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
