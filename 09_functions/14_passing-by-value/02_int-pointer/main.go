package main

import "fmt"

func main() {
	age := 44
	fmt.Println("main age pointer: ", &age) // age pointer
	changeMe(&age)                          // &

	fmt.Println("main: ", age)
}

func changeMe(z *int) { // *int
	fmt.Println("1st: in function:  ", z)
	*z = 34
	fmt.Println("2nd: in function: ", z)
}
