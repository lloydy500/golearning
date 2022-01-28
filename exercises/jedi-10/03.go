package main

import (
	"fmt"
)

// pull the values of the channel using a for range loop
func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// put values 0-99 onto channel and rtn channel

func gen() <-chan int {
	c := make(chan int)

	go func() { // starts a goroutine; adds values to c
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // closes channel
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
