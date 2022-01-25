package main

import (
	"fmt"
	"math"
)

func main() {
	sq := square{
		L: 14,
	}
	ci := circle{
		R: 42,
	}
	info(sq)
	fmt.Printf("%T\n", &ci)
	info(&ci)
}

type square struct {
	L float64
}

type circle struct {
	R float64
}

// method sets determine what methods attach to a TYPE.
// POINTER method set
// works with VALUES that are POINTERS.

// NON-POINTER method set
// works with VALUES that are POINTERS or NON-POINTERS
func (s square) area() float64 {
	return s.L * s.L
}
func (c *circle) area() float64 {
	return math.Pi * c.R * c.R
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
