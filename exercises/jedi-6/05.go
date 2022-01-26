package main

import (
	"fmt"
	"math"
)

// create  a type square
// create a type circle
// attach a mehtod to each that calculates area and returns it
// create a type shape which defines an
// interface as anything which has  the area method
// create a value of type square
// create a value of type circle
// use func info to print out the area of square
// use func info to print out the area of circle

type square struct {
	L float64
}

type circle struct {
	R float64
}

func (s square) area() float64 {
	return s.L * s.L
}
func (c circle) area() float64 {
	return math.Pi * c.R * c.R
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		L: 14,
	}
	ci := circle{
		R: 42,
	}
	info(sq)
	info(ci)
}
