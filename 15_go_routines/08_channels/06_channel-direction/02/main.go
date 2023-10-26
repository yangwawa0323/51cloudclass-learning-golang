package main

import "fmt"

func main() {
	// c := incrementor() // return chan int
	// cSum := puller(incrementor()) // return chan int
	// for n := range cSum {
	for n := range puller(incrementor()) {
		fmt.Println("Sum : ", n)
	}
}

func incrementor() <-chan int {
	out := make(chan int) // unbufferred
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// c chan `0`
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c { // <-c
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
