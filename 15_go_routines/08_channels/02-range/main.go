package main

import (
	"fmt"
)

func main() {

	c := make(chan int) // unbufferred

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending: ", i)
			c <- i
		}
		close(c)
		// if channel closed , receive channel remain bool set to false

	}() // IIFE

	//   num , remain := <- c
	for n := range c {
		fmt.Println("Received :", n)
	}
}
