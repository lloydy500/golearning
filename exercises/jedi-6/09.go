package main

import "fmt"

// a callback is when we pass a func into afunc
//  pass a func into a func as an argument

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := accumulator(ii...)
	fmt.Println(a)
	b := five_multiple(accumulator, ii...)
	fmt.Println(b)
}

func accumulator(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func five_multiple(f func(i ...int) int, xii ...int) int {
	var vii []int
	for i, v := range xii {
		if v%5 == 0 {
			fmt.Println(i, v)
			vii = append(vii, 999)

		} else {
			vii = append(vii, v)
		}
	}
	return f(vii...)
}
