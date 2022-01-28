package main

import "fmt"

// write a program that launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for v := 0; v < 10; v++ {
				c <- v
				fmt.Printf("goroutine %v: \t added %v to the channel\n", i, v)
			}
			close(c)
		}()
	}

	for val := range c {
		fmt.Printf("took %v off the channel\n", val)

	}

}
