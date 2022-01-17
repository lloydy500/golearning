package main

import "fmt"

func main() {

	x := 42 // declare the variable with short declaration operator
	fmt.Println(x)
	x = 99 // don't need to declate to re assign
	fmt.Println(x)
	y := 100 + 21  // (100 PLUS 21 IS EXPRESSION, AND THE WHOLE LINE IS A STATEMENT)
	fmt.Println(y) // must use any value declared
}

// func main() {
// 	n, _ := fmt.Println("hello world")
// 	fmt.Println(n)
// }

// func main() {
// 	fmt.Println("hello world!")
// 	foo()
// 	fmt.Println("Something else")

// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// 	bar()
// }

// func foo() {
// 	fmt.Println("I'm in foo")
// }

// func bar() {
// 	fmt.Println("and we're all done!", true, 23.4)
// }
