package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // lock this variable whilst writing
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // then unlock for use by another  goroutine
			wg.Done()
		}()
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
