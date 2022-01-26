package main

import (
	"fmt"
	"sort"
)

// type SortByInt []int

// func (a SortByInt) Len() int           { return len(a) }
// func (a SortByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortByInt) Less(i, j int) bool { return a[i] < a[j] }

// type SortByString []string

// func (s SortByString) Len() int           { return len(s) }
// func (s SortByString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s SortByString) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

}
