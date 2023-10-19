package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)
	var b = &a
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
}
