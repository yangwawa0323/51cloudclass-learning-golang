package main

import "fmt"

func main() {
	c := make(chan int, 1) // bufferred channel

	c <- 42
	c <- 43 // send int to channel, block. reason : size

	fmt.Println(<-c)
}
