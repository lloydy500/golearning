package main

import "fmt"

func main() {
	x := 40
	if x == 40 {
		fmt.Println("our value is 40")
	} else if x == 41 {
		fmt.Println("our value is 41")
	} else {
		fmt.Println("our value is not 40 or 41")
	}
}
