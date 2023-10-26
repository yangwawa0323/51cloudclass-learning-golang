package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}

func incrementor(lable string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		x++
		counter = x
		fmt.Println(lable, i, "Counter: ", counter)
	}
	wg.Done() // counter decrement
}
