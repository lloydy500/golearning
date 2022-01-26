package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Printf("My name is %v, I'm %v years old.\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Sandy",
		14,
	}
	p1.speak()
	saySomething(&p1)
	// saySomething(p1)

}

// type circle struct {
// 	radius float64
// }

// type shape interface {
// 	area() float64
// }

// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func info(s shape) {
// 	fmt.Println("area:", s.area())
// }
// func main() {
// 	c := circle{5}
// 	// info(c)
// 	// info(&c)
// 	fmt.Println(c.area())
// }
