package main

import "fmt"

// print out the remainder (modulus) which is found for each number
// between 10 and 100 when it's divided by 4

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("when %v is divided by 4, the remainder is %v\n", i, i%4)
	}
}
