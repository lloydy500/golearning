package main

import "fmt"

// create a program that uses a SWITCH STATEMENT with the SWITCH
// EXPRESSION specified as a variable of TYPE string with the IDENTIFIER "favSport"
func main() {
	var favSport string = "snowboarding"

	switch favSport {
	case "snowboarding":
		fmt.Println("rad...")
	case "football":
		fmt.Println("wat lad?")
	case "squash":
		fmt.Println("get outside")
	}
}
