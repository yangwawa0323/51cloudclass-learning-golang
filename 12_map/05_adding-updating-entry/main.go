package main

import "fmt"

func main() {
	var myGreeting = map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour.",
	}
	fmt.Println(myGreeting)

	myGreeting["Tim"] = "你好"        // replace
	myGreeting["Harleen"] = "Howdy" // insert
	fmt.Println(myGreeting)
}
