package main

import "fmt"

var x int // default: 0

func main() {
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
	fmt.Println(increment())
}

func increment() int {
	x++
	return x
}
