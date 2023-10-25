package main

import "fmt"

func main() {
	name := "Yangwawa"
	str, ok := name.(string) // assertion   👎
	if ok {
		fmt.Println("%q\n", str) // ""
	} else {
		fmt.Println("value is not a string\n")
	}
}
