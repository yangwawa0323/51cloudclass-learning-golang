package main

import "fmt"

func main() {
	c := make(chan struct{ num int })

	go func() {
		for i := 0; i < 10; i++ {
			c <- struct{ num int }{i}
		}
	}() // block in goroutines

	fmt.Println(<-c)
	fmt.Println("Bye bye")
}
