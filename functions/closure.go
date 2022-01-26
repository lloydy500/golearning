package main

import "fmt"

var x int //global scope

func main() {

	fmt.Println(x)
	x++
	// you can make a code block inside a code block
	{
		y := 42
		fmt.Println(y)

	}
	foo()
	fmt.Println(x)

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func foo() {
	fmt.Println("adding one to x...")
	x++
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
