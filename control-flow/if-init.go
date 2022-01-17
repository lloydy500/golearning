package main

import "fmt"

func main() {
	// go has semicolons at the end of every statment
	// most of the time we don't have to put then in
	// becausae the compiler does it for us
	// unless we want to make more than one statement on 1 line
	if x := 42; x == 2 {
		fmt.Println("001")
	}
}
