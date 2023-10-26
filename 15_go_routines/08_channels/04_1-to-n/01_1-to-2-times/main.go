package main

import "fmt"

func main() {
	c := make(chan int)     // unbufferred
	done := make(chan bool) // semaphore

	// Send data goroutine
	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}() // IIFE

	// 1/2 Receive goroutine
	go func(s string) {
		for n := range c {
			fmt.Println(s, "sending -> ", n)
		}
		done <- true
	}("thread 0")

	// 2/2 Receive goroutine
	go func(s string) {
		for n := range c {
			fmt.Println(s, "sending     -> ", n)
		}
		done <- true
	}("thread 1")

	<-done // blocked
	<-done // blocked
}
