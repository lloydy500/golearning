package main

import "fmt"

// create a new type, 'vehicle'
// the underlying type is struct
// add doors and color as fields

type vehicle struct {
	doors int
	color string
}

// create two new types: truck and sedan.
// the underlying type of each of these new types is a struct.
// embed the "vehicle" type in both truck and sedan.
// give truck the field fourWheel which will be set to bool.
type truck struct {
	vehicle
	fourWheel bool
}

// give sedan the field luxury which will be set to bool
type sedan struct {
	vehicle
	luxury bool
}

// using the above structs:
// using a COMPOSITE LITERAL, create a value of type truck and
// assign values to the fields;
// using a composite literal, create a value of type sedan and
// assign values to the fields
func main() {
	v1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}
	v2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "orange",
		},
		luxury: false,
	}
	// print out each of these values
	// print out a single field from each of these values.
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.color)
	fmt.Println(v2.luxury)

}
