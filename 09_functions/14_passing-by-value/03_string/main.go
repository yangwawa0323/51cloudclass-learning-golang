package main

import "fmt"

func main() {
	name := "yangwawa"

	fmt.Println(name)

	changeMe(name)

	fmt.Println(name)
}

func changeMe(z string) {
	fmt.Println("changeMe func : ", z)
	z = "Jake"
	fmt.Println("changeMe func : ", z)
}
