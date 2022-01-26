package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type SortByAge []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type SortByLast []user

func (s SortByLast) Len() int           { return len(s) }
func (s SortByLast) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByLast) Less(i, j int) bool { return s[i].Last < s[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// your code goes here
	// sort the []user by:
	// age
	sort.Sort(SortByAge(users))
	fmt.Println(users)
	// last
	sort.Sort(SortByLast(users))
	fmt.Println(users)

	// Then sort each [] Sayings for each user
	// print everything out using range

	// remember sort change the underlying ARRAY
	for _, v := range users {
		fmt.Println(v.First, v.Last, ":", v.Age)
		fmt.Println("\tThe user's favourite sayings are: ")
		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}

	}
}
