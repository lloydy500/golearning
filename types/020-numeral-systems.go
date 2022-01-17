package main

import "fmt"

func main() {
	//letter h
	s := "H"
	fmt.Println(s)
	// byte slice
	bs := []byte(s)
	fmt.Println(bs)
	// what is the TYPE of the item at position 0?
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	// convert to binary (base2)
	fmt.Printf("%b\n", n)
	// convert to hexadecimal (base16)
	fmt.Printf("%#X\n", n)

}
