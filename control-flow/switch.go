package main

import "fmt"

func main() {
	switch "Bond" {
	case "Moneypenny", "Bond", "Dr. No":
		fmt.Println("Money, Bond or No")
	case "Q":
		fmt.Println("This is Q")
	}
}

// func main() {
// 	switch {
// 	case false:
// 		fmt.Println("shouldn't print")
// 	case 2 == 4:
// 		fmt.Println("also shouldn't print")
// 	case 3 == 3:
// 		fmt.Println("should print")
// 		fallthrough // let the compiler fallthrough to the next case
// 		// strangely, fallthrough will run all print statements after it
// 		// avoid using it
// 	case 4 == 4:
// 		fmt.Println("also true but does it print?")
// 	}
// }
