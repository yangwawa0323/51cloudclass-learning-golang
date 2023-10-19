package main

import "fmt"

// global_message := "Hi!"  👎

var global_message string = "Hi!" // 👍

func main() {

	// shorthand is only inside function
	message := "Hello world."
	a, b, c := 1, false, 3
	d := 4
	e := true

	fmt.Println(global_message, message, a, b, c, d, e)
}
