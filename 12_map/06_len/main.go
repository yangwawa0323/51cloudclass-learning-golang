package main

import "fmt"

func main() {
	var myGreeting = map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour.",
	}
	fmt.Println(myGreeting)

	myGreeting["Harleen"] = "Howdy" // insert
	fmt.Println(len(myGreeting))
}
