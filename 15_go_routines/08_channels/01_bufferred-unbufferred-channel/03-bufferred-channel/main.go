package main

import "fmt"

func main() {
	c := make(chan int, 1) // bufferred channel

	c <- 42 // send int to channel, block

	fmt.Println(<-c)
}
