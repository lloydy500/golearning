package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	// now bar sits within the main function instead of inside goroutine
	// therefore the main function waits for bar to unblock
	// the channel, c, before it can complete.
	// so we don't need to use a waitgroup here.
	bar(c)
	fmt.Println("about to exit...")
}

// notice that we defined c as a bidirectional channel in main()
// then we use funcs to specifiy send or recieve

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
