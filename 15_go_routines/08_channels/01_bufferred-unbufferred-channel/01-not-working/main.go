package main

import "fmt"

func main() {
	c := make(chan int) // unbufferred channel

	c <- 42 // send int to channel, block

	fmt.Println(<-c)
}
