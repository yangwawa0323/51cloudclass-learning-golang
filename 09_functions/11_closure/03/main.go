package main

import "fmt"

var x = 100

func main() {
	// x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
