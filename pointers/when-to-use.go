package main

import "fmt"

// everything in Go is pass by value!

// step 1 no pointer

// func main() {
// 	x := 0
// 	foo(x)
// 	fmt.Println(x)
// }

// func foo(y int) {
// 	fmt.Println(y)
// 	y = 43
// 	fmt.Println(y)

// }

// step 2 with pointer

func main() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x) // address
	fmt.Println("x after", x)
	fmt.Println("x after", &x)

}

func foo(y *int) { //function takes some address
	fmt.Println("y befor", y)  // address
	fmt.Println("y befor", *y) // dereferenced address
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)

}
