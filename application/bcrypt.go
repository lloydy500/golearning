package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// we never store passwords, only the encrypted versions
// of those passwords. from the docs:
// func GenerateFromPassword(password []byte, cost int) ([]byte, error)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
}
