package main

import (
	"fmt"
)

func main() {
	c := make(chan int)     // data channel
	done := make(chan bool) // signal/semaphore channel

	go func(s string) {

		for i := 0; i < 10; i++ {
			fmt.Println(s, ": sending  ", i)
			c <- i
		}
		done <- true // send bool to signal channel
	}("thread 1")

	go func(s string) {
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println(s, ": sending  ", i)
			c <- i
		}
		done <- true
	}("thread 2")

	// From a channel receive data, if has no send data , it is block
	// t1 : <- done   blocked
	go func() {
		<-done // t1. receive, blocked
		<-done // t2. receive , blocked
		close(c)
	}()

	for n := range c {
		fmt.Println("Received :", n)
	}
}
