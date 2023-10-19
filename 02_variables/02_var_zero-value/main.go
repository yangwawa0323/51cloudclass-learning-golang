package main

import "fmt"

func main() {
	var a int
	var b string
	var c float32
	var d bool

	fmt.Printf("%v \n", a)     //  default : 0
	fmt.Printf("\"%v\" \n", b) //  default : ""
	fmt.Printf("%v \n", c)     //  default : 0
	fmt.Printf("%v \n", d)     //  default : false

	// newly defined a variable with default type according
	// variable type assertion
	d = true
	fmt.Printf("%v \n", d) //  default : false
}
