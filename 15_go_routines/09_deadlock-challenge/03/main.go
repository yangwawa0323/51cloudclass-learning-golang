package main

import "fmt"

func main() {
	c := make(chan struct{ num int })

	go func() {
		for i := 0; i < 10; i++ {
			c <- struct{ num int }{i}
		}
	}() // block in goroutines

	for {
		fmt.Println(<-c) // block no more data received
	}
	fmt.Println("Bye bye")
}
