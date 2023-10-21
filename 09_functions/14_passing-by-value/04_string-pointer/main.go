package main

import "fmt"

func main() {
	name := "yangwawa"

	fmt.Println(name)
	fmt.Println("main name is : ", &name)

	changeMe(&name)

	fmt.Println(name)
}

func changeMe(z *string) { // pointer type string
	fmt.Println("changeMe z pointer : ", z)

	*z = "Jake"
	fmt.Println("changeMe func : ", z)
}
