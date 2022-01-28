// get this code working:

package main

import "fmt"

func main() {

	// ○ a buffered channel
	// this is not an ideal solution
	// ■ solution: https://play.golang.org/p/Y0Hx6IZc3U
	c := make(chan int, 1)

	c <- 42
	fmt.Println(<-c)

}
