package main

import "fmt"

func main() {
	var num = 7

	if num%2 == 0 {
		fmt.Printf("%d is even 👍\n", num) // 👍
	} else {
		fmt.Printf("%d is odd 👎\n", num) // 👎
	}
}
