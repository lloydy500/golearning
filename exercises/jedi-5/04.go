package main

import "fmt"

// create and use an anonymous data struct
// challenge: one field of data type map, another field of type slice

func main() {
	a := struct {
		shopping_basket map[string]int
		popular_items   []string
	}{
		shopping_basket: map[string]int{
			"apples":      2,
			"bananas":     4,
			"light bulbs": 1,
		},
		popular_items: []string{"pasta", "loo roll", "tinned beans"},
	}
	fmt.Println("I picked up:")
	for k, v := range a.shopping_basket {
		fmt.Printf("%v %v\n", v, k)
	}
	fmt.Println("I also noticed the following items were popular:")
	for i, val := range a.popular_items {
		fmt.Println(i, val)
	}

}
