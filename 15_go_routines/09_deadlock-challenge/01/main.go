package main

import "fmt"

func main() {
	c := make(chan int)
	// go func() {
	c <- 1 // block
	// }()
	fmt.Println(<-c)
}
