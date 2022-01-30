package main

// fix using mutex

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0
	gs := 200
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			// the race condition is where we access x
			// so this is where we start the lock
			mu.Lock()
			x := total
			// runtime.Gosched() it makes sense to remove this
			x++
			total = x
			fmt.Println(total) // this is a read so keep in the lock
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(total)
	fmt.Println("Finished")
}
