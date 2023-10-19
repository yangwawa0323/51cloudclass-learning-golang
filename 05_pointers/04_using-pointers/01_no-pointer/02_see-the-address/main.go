package main

import "fmt"

func zero(val int) { // val = 50 only exists zero function
	fmt.Printf("%p\n", &val) // %p pointer
	fmt.Println(&val)
	val = 0
}

func main() {
	x := 50
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)        // x = 50
	fmt.Println(x) // x = 0 ?
}
