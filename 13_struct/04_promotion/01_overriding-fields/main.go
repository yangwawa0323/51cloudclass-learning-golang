package main

import "fmt"

// struct fields , miss field use default , age int default is 0
type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person        person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bouds",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If look could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.person.First, p1.First)
	fmt.Println(p2.person.First, p2.First)
}
