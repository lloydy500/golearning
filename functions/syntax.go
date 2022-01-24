package main

import "fmt"

// func (r receiver) identifier(parameters (return(s)) { ... })
func main() {
	foo()
}
func foo() {
	fmt.Println("hello from foo")
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

}

// we define a func with PARAMETERS (if any)
// we call a func with ARGUMENTS (if any)

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := true
	return a, b
}

// everything in Go is PASS BY VALUE
