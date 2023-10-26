package main

import "fmt"

func main() {

	n := 100000
	c := make(chan int)     // unbufferred data
	done := make(chan bool) // semaphore

	// Send data goroutine
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}() // IIFE

	for i := 0; i < n; i++ {
		// i/n Receive goroutine
		go func(s string) {
			for n := range c {
				fmt.Println(s, "sending -> ", n)
			}
			done <- true
		}(fmt.Sprintf("thread %d", i))
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
