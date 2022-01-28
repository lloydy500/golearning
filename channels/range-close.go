package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	// recieve
	// range will pull off the channel until its closed
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit...")
}

// notice that we defined c as a bidirectional channel in main()
// then we use funcs to specifiy send or recieve
