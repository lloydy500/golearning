package main

import "fmt"

func main() {
	// we don't really use arrays in go
	// unless we're going something really specific
	var x [5]int // arrays have to be same TYPE and length is part of TYPE
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

}
