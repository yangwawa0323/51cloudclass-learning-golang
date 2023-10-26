package main

import (
	"fmt"
)

func main() {
	n := 500
	c := make(chan int)     // unbufferred data channel
	done := make(chan bool) // signal/semaphore channel

	for i := 0; i < n; i++ {
		go func(s string) {

			for i := 0; i < 10; i++ {
				fmt.Println(s, ": sending  ", i)
				c <- i
			}
			done <- true // send bool to signal channel
		}(fmt.Sprintf("thread %d", i))
	}

	// From a channel receive data, if has no send data , it is block
	// t1 : <- done   blocked
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println("Received :", n)
	}
}
