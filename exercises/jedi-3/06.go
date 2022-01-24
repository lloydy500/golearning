package main

// create a program that shows the 'if'
// statement in action

import "fmt"

func main() {
	today := "Lovely sunshine"
	if today == "thunder and lightning" {
		fmt.Println("I'm going to have to wear clothes")
	} else if today == "snowy" {
		fmt.Println("I've got to at least wear something")
	} else {
		fmt.Println("I'm going out... naked")
	}
}
