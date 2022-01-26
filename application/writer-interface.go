package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hi from Println")
	fmt.Fprintln(os.Stdout, "Hi from Fprintln")
	io.WriteString(os.Stdout, "Hi from WriteString")
}
