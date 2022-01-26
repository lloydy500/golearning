package main

import "fmt"

// closure is when we have "enclosed" the scope of a
// variable in some code block. Create a func which "encloses"
// the scope of a variable

func main() {
	x := doubler()   // the scope of z is enclosed in these variable and stored in mem for each
	y := doubler()   // the scope of z is enclosed in these variable and stored in mem for each
	fmt.Println(x()) // 4
	fmt.Println(x()) // 8
	fmt.Println(x()) // 16
	fmt.Println(x()) // 32
	fmt.Println(x()) // 64
	fmt.Println(y()) // 4 z is initialised and stored at a different place in memory
	fmt.Println(y()) // 8
	fmt.Println(x()) // 128

}

func doubler() func() int {
	fmt.Println("Doubling x...")
	z := 2
	return func() int {
		z *= 2
		return z
	}
}
