package main

import "fmt"

// create a for loop using the syntax
// for {}
// print all the years you've been alive
func main() {
	bd := 1992
	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
