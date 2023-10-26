package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 30; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring; I'm go")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // sender data coming from receiver channel
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(name string) <-chan string {
	c := make(chan string) // unbufferred channel
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
