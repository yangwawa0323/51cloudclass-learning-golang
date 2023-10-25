package main

import "fmt"

func main() {
	var x = 12
	var y = 12.72121212
	// z := x + int(y)  // float64 to int
	z := float64(x) + y // int to float64
	fmt.Printf("%T: %v\n", z, z)

	message := []byte("Hello")
	fmt.Printf("%T: %v\n", message, message)
	message2 := string(message)
	fmt.Printf("%T: %v\n", message2, message2)
}
