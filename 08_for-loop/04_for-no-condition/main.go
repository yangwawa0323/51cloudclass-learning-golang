package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0 // initial
	for {  // condition ?  endless loop
		fmt.Println(i)
		time.Sleep(time.Second) // Waiting for new connection, block
		i++
	}
}
