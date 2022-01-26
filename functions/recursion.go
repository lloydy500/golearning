package main

import "fmt"

// recursion is when a func calls iself
// a certain number of times then stops

func main() {
	n := factorial(4)
	fmt.Println(n)
	m := loop_factorial(4)
	fmt.Println(m)
}

// 4 * 3* 2 * 1 * 1
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// challenge: write this code using a loop

func loop_factorial(i int) int {
	total := 1
	for ; i > 0; i-- {
		total *= i
	}
	return total
}

//  my answer ( not quite as neat )
// func loop_factorial(i int) int {
// 	total := 0
// 	for x := i - 1; x > 0; x-- {
// 		total += x * i
// 	}
// 	return total
// }
