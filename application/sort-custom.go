package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 36}

	people := []Person{p1, p2, p3, p4}
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
