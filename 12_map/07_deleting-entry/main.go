package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":  "Good morning",
		"one":   "Bonjour!",
		"two":   "你好",
		"three": "Bongiorno!",
	}
	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println("two is : ", myGreeting["two"])
	fmt.Println(myGreeting)
}
