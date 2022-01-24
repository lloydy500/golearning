package main

import "fmt"

// create a program that uses the switch statement
// with no SWITCH EXPRESSION specified

func main() {
	switch {
	case 2 == 3:
		fmt.Println("I never knew that!")
	case true && false:
		fmt.Println("Surely not?")
	case !false:
		fmt.Println("Print me and break out of branch!")
	case 1 < 2:
		fmt.Println("This won't run")
	}
}
