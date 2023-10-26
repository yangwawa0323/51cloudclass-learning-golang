package main

import "fmt"

func main() { // main go routine
	c := make(chan int) // unbufferred channel

	go func() { // go func1 routine
		c <- 42 // send int to channel, block
	}() // IIFE

	fmt.Println(<-c) // receive channel
}
