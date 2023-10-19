package main

import "fmt"

func zero(val int) { // val = 50
	val = 0
}

func main() {
	x := 50
	zero(x)        // x = 50
	fmt.Println(x) // x = 0 ?
}
