package main

import "fmt"

func main() {
	greet("Yang", "wawa")
	greet("John", "Doe")
	hello("Jakie", "Chen")
}

func greet(fname string, lname string) {
	fmt.Println("Hi! ", lname, fname)
}

func hello(fname, lname string) {

	fmt.Println("Hi! ", fname, lname)
}
