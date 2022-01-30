package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	x[0] = 42
	x[9] = 1231
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// can't do x[10] beyond limit
	// but we can append up until capacity
	x = append(x, 1024)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 1025)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// here we exceed the cap of the underlying array
	// the runtime creates a new array of double capacity and deletes the old one
	// this is slow and takes computational power
	x = append(x, 1026)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
