package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := 'M'

	fmt.Printf("%T \n", a) // %v verb --> Stringer
	fmt.Printf("%T \n", b) // %T type
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}
