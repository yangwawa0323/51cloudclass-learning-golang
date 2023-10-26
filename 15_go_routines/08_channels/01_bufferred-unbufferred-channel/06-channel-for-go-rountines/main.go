package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup // 👎

func main() {

	c := make(chan int) // unbufferred

	// wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending: ", i)
			c <- i
		}
		// wg.Done()
	}() // IIFE

	go func() {
		for {
			fmt.Println("Received: ", <-c)
		}
		// wg.Done()  // 👎 unreachable code
	}()

	// wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("Bye bye")
}
