package main

import "fmt"

func main() {
	var x [58]string

	// Ascii
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
}
