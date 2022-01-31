package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-12.23)
	if err != nil {
		log.Println(err)
	}
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{lat: "487221 E", long: "487221 E", err: fmt.Errorf("this is the error: %v", f)}
	}
	return 42, nil
}
