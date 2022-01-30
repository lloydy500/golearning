package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Stawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
