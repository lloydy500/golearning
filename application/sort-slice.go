package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{5, 2, 7, 9, 3, 6, 3, 1, 67, 0, 2, 3}
	xs := []string{"hello", "this", "is", "a", "slice"}
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("--------------")
	sort.Strings(xs)
	fmt.Println(xs)

}
