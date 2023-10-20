package main

import "fmt"

func main() {
	if !true { // false
		fmt.Println("This is running")
	}

	if !false { // true
		fmt.Println("This is never show")
	}
}
