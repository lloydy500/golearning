package main

import "fmt"

// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns and int and a string
// call both func
// print out their results

func main() {
	i := foo()
	x, s := bar()

	fmt.Println(i, x, s)
}

func foo() int {
	return 1984
}

func bar() (int, string) {
	return 007, "James Bond"
}
