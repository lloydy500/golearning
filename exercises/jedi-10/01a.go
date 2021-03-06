// get this code working:
// ○ a buffered channel
// ■ solution: https://play.golang.org/p/Y0Hx6IZc3U

package main

import "fmt"

func main() {

	// ○ using func literal, aka, anonymous self-executing func
	// ■ solution: https://play.golang.org/p/SHr3lpX4so
	c := make(chan int)

	go func() {

		c <- 42
	}()

	fmt.Println(<-c)

}
