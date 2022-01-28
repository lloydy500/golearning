package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c: // first try to get value off c
			fmt.Println("from channel c:", v)
		case <-q: // just return from func if we get here
			return
		}
	}
}
