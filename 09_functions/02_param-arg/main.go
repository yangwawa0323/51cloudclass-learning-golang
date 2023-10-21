package main

import "fmt"

func main() {
	greet("Yangwawa")
	greet("John")
	greet("Jake")
}

func greet(name string) {
	fmt.Println("Hi! ", name)
}
