package main

// fix using mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var total int64
	gs := 200
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&total, 1)
			// runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&total)) // this is a read so keep in the lock
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(total)
	fmt.Println("Finished")
}
