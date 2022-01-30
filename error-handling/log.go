package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println("err happened", err)
		// log.Fatalln(err) // same a log and then calling os.Exit(1), deferred funcs won't run
		// log.Panicln(err)
		panic(err)
	}
	defer f2.Close()

	fmt.Println("Check log file in this directory")
}
