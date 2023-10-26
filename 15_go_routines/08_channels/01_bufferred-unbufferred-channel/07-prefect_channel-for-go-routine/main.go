package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int) // unbufferred

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending: ", i)
			c <- i
		}
		close(c)
		// if channel closed , receive channel remain bool set to false
		wg.Done()
	}() // IIFE

	go func() {
		for {
			num, remain := <-c // receive channel
			if !remain {
				fmt.Println("Received all data")
				break
			}
			fmt.Println("Received: ", num)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Bye bye")
}
