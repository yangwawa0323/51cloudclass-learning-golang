package main

import "fmt"

func main() {
	childMath := doMath(4)
	fmt.Printf("%T\n", childMath)
	fmt.Println(childMath(10, 9))
}

func doMath(level int) func(a, b int) int {
	add := func(a, b int) int {
		return a + b
	}

	mul := func(a, b int) int {
		return a * b
	}

	if level <= 2 {
		return add

	} else {
		return mul
	}
}
