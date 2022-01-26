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
	for i := 0; i < gs; i++ {
		go func() {
			// each go routine will access a different variable, v
			// becuase v's scope is limited to this FUNCTION LITERAL
			// counter is outside of this scope so can be accessed
			// by each goroutine
			v := counter
			// time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
