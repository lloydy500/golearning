package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}

	_, err = os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}

}
