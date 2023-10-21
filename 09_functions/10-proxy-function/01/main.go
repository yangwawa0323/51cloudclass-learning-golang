// proxy

package main

import "fmt"

func main() {
	greet := makeGreeter() // func() string
	fmt.Println(greet())
}

func makeGreeter() func() string { // string , bool , int ,
	greet := func() string {
		return "Hello world!"
	}
	return greet
}
