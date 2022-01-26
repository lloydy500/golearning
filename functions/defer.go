package main

import "fmt"

// when func main() exists, defer statement is run

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
