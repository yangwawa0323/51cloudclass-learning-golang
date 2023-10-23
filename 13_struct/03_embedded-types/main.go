package main

import "fmt"

// struct fields , miss field use default , age int default is 0
type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person // == person person
	// while a field only has name/type, all the struct's attributes
	// is transform to current struct
	// First string
	// Last  string
	// Age   int
	/////////////////////
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bouds",
			Age:   20,
		},
		LicenseToKill: true,
	}
	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last)
	fmt.Println(p2.Last, p2.Age)
}
