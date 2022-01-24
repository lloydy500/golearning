package main

import "fmt"

// using a for loop print out
// the years you've been alive
// use syntax:
// for [condition] {}

func main() {
	bd := 1992
	for bd <= 2022 {
		fmt.Println(bd)
		bd++
	}

}
