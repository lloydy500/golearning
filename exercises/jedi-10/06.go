package main

import "fmt"

// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

func main() {
	c := make(chan int)
	go fill(c)
	empty(c)
	fmt.Println("exiting main")

}

func fill(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Printf("added %v to channel c\n", i)
	}
	close(c)
}

func empty(c chan int) {
	for v := range c {
		fmt.Printf("taking %v from channel c\n", v)
	}
}
