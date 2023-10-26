package main

import "fmt"

func main() {
	c := make(chan struct{ num int })

	go func() {
		for i := 0; i < 10; i++ {
			c <- struct{ num int }{i}
		}
		// close(c)     // IMPORTANT!!! IMPORTANT!!! IMPORTANT!!!
	}() // block in goroutines

	for s := range c { // range iterate end at closed singal
		fmt.Println(s) // block no more data received
	}
	fmt.Println("Bye bye")
}
