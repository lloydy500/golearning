package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// create a custom error message using fmt.Errorf

// could I use errors.New? or log.Println, log.Panic, log.Fatal?

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred.", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

//toJSON needs to rtn an error and byteslice
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("Error in toJSON: %v", err)) // we can pass in errors new
		// errors.New takes a string returns an error
	}
	return bs, nil
}
