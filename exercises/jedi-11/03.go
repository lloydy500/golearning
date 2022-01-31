package main

import "fmt"

// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
// has a value of type error as a parameter. Create a value of type “customErr” and pass it into
// “foo”. If you need a hint, here is one.
func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1) // notice we can pass in c1 because it is of type customErr and error
}

type customErr struct {
	info string
}

func (ce customErr) Error() string { // custom error implements the method Error, so any customErr is also of type error
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println("Foo ran: ", e, "\n", e.(customErr).info) // this is ASSERTION
	// we can't call e.info because e is TYPE error, not customErr
	// we need to ASSERT e's type for the compiler.
}
