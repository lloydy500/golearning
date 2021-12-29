package main

import "fmt"

//using iota, create 4 CONSTANTS for the last 4 years.
// Print the constant values.

const (
	y1 = 2021 - iota
	// the y1 statement is implied for all other y values
	y2
	y3
	y4
)

func main() {
	fmt.Println("This year is", y1)
	fmt.Println("The years before were", y2, y3, y4)

}
