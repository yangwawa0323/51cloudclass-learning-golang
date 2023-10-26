package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//
	wg.Add(2) // counter = 2 go routines
	go foo()  //
	go bar()  //
	// waiting
	wg.Wait() // block until counter == 0
	fmt.Println("Bye bye")
}

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done() // counter - 1
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(5 * time.Millisecond)
	}
	wg.Done() // counter - 1
}
