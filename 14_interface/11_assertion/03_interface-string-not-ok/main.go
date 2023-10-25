package main

import "fmt"

func main() {
	var name interface{} = 7
	str, ok := name.(string) // assertion
	if ok {
		fmt.Printf("%q %T\n", str, str) // ""
	} else {
		fmt.Printf("value is not a string\n")
	}
}
