package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	total := 0
	gs := 200
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			x := total
			runtime.Gosched()
			x++
			total = x
			fmt.Println(total)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(total)
	fmt.Println("Finished")
}
