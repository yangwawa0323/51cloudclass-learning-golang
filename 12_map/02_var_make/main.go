package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	// 1st string key
	// 2nd string value
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Tim"])
}
