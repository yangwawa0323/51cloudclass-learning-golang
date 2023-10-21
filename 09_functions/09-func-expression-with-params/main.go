package main

import "fmt"

// javascript  ()=> {}
// func greeting(name string) {}

func main() {
	greeting := func(name string) { //anonymous
		fmt.Println("H! ", name)
	}

	fmt.Printf("%T\n", greeting)
	greeting("yangwawa")

}
