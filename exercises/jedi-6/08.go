package main

import "fmt"

// create a func which returns a func
// assign the returned func to a variable
// call the returned func

func main() {
	duck := quack_caller()
	s := duck()
	fmt.Println(s)
}

func quack_caller() func() string {
	return func() string {
		return "Quack!!"
	}
}
