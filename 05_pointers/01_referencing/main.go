package main

import "fmt"

func main() {
	a := 45
	// var b = true

	fmt.Println(a)
	fmt.Println(&a)
	// fmt.Printf("%T, %T \n", a, &a)
	// fmt.Printf("%T, %T \n", b, &b)

	var c = &a
	fmt.Println(c)
}
