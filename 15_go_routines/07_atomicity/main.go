package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

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

		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(lable, i, "Counter: ", atomic.LoadInt64(&counter))
	}
	wg.Done() // counter decrement
}
