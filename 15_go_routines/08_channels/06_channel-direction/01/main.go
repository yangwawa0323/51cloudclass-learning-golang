package main

import "fmt"

func main() {
	c1 := make(chan int, 1)

	c1 <- 10

	receive_only_channel := make(<-chan int, 1)
	receive_only_channel <- 2 // 👎 cannot send to receive-only channel receive_only_channel

	send_only_channel := make(chan<- int, 1)
	// send_only_channel <- 100
	<-send_only_channel // 👎 cannot receive from send-only channel send_only_channel
	fmt.Println("Bye bye")
}
