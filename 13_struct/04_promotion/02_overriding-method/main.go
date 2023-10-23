package main

import "fmt"

// struct fields , miss field use default , age int default is 0
type person struct {
	Name string
	Age  int
}

func (p person) Greeting() {
	fmt.Println("I am just a regular person")
}

type doubleZero struct {
	person        person
	LicenseToKill bool
}

func (doubleZero) Greeting() {
	fmt.Println("Hi! 00")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James bouds",
		},
		LicenseToKill: true,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	p1.Greeting()
	p2.Greeting()
	fmt.Printf("%T\n", p2.person)
	p2.person.Greeting()
}
