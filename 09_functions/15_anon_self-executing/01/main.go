// Javascript IIFE
// Immediately Invoked Function Expression

package main

import "fmt"

func main() {
	func() {
		// defined variables in closure
		fmt.Println("Call immediately")
	}()

	// variable
}
