package main

import (
	"fmt"
	"runtime"
	"sync"
)

// in addition to the main goroutine
// launch two additional goroutines
// each additional goroutine should print something out
// use  waitgroups to make sure each goroutine finishes before
// your program exists

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	fmt.Println(runtime.NumGoroutine())
	bar()
	fmt.Println("waiting for goroutines")
	wg.Wait()
	fmt.Println("Finished.")
}

func foo() {
	fmt.Println("Hello from foo")
	wg.Done()
}

func bar() {
	fmt.Println("Hello from bar")
}
