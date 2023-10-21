package main

import "fmt"

// javascript  ()=> {}

func main() {
	greeting := func() { //anonymous
		fmt.Println("Hello world")
	}

	fmt.Printf("%T\n", greeting)
	greeting()

}
