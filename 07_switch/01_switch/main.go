package main

import "fmt"

func main() {
	var name = "John"
	switch name {
	case "Daniel":
		fmt.Println("Nice to meet you,", name)
	case "Yangwawa":
		fmt.Println("It is long time no see", name)
	case "Jenny":
		fmt.Println("Hi, Jeney")
		// default:
		// fmt.Println("Does we meet before?")
	}

	// no need break
	fmt.Println("Bye bye!")
}
