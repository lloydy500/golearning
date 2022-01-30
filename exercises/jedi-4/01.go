package main

// using a COMPOSITE LITERAL
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES TO EACH INDEX POSITION
import "fmt"

func main() {
	var x [5]int // the size of the array is part of the TYPE
	x[0] = 1
	x[1] = 23
	x[2] = 42
	x[3] = 25
	x[4] = 102

	for k, v := range x {
		fmt.Println(k, v)
	}
	fmt.Printf("%T\n", x)
}
