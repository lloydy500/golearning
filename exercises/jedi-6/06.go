package main

import "fmt"

func main() {
	five_square := func(i int) int {
		return i * i
	}(5)
	fmt.Println(five_square)
}
