package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)

	var b = &a // b type   ---> *int
	fmt.Println(b)

	*b = 42
	fmt.Println(b)
	fmt.Println(a)
}
