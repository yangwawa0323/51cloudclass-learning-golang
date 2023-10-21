package main

import (
	"fmt"
	"strings"
)

func main() {
	// message := greet("Yang", "wawa")
	// fmt.Println(message)
	people := "Liusan"

	// findLiu(string) -> bool    toLowerName(string) -> string
	// findLiu(toLowerName(string) -> string)  -> bool
	// findLiu [[ toLowerName ]] --> bool
	////////////////////////////////////////////
	// people = toLowerName(people)
	// findLiu(people)
	if findLiu(toLowerName(people)) {
		fmt.Println("Long time no see!")
	}
	fmt.Println("Bye bye!")
}

// func greet(fname, lname string) string {
// 	return fmt.Sprintf("Hi! %s %s", fname, lname)
// }

func findLiu(name string) bool {
	return strings.HasPrefix(name, "liu")
}

func toLowerName(name string) string {
	return strings.ToLower(name)
}
