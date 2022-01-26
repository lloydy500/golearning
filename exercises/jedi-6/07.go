package main

import "fmt"

var g func()

func main() {
	f := func() {
		for i := 0; i < 10; i++ {
			switch {
			case i%2 == 0:
				fmt.Println("even")
			case i%2 != 0:
				fmt.Println("odd")
			}
		}
	}
	g = f
	g()
}
