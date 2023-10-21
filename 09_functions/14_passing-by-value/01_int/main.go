package main

import "fmt"

func main() {
	age := 44
	changeMe(age)

	fmt.Println("main: ", age)
}

func changeMe(z int) {
	fmt.Println("1st: in function: ", z)
	z = 34
	fmt.Println("2nd: in function: ", z)
}
