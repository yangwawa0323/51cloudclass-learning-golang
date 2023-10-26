package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	go func(s string) {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println(s, ": sending  ", i)
			c <- i
		}
		wg.Done()
	}("thread 1")

	go func(s string) {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println(s, ": sending  ", i)
			c <- i
		}
		wg.Done()
	}("thread 2")

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println("Received :", n)
	}
}
