package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 6, 7) // receiver channel
	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n) // 9, 4 , 36 ,25, 49
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// cs type  []<-chan int
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
