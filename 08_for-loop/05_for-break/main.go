package main

import "fmt"

func main() {
	for j := 0; j < 3; j++ {
		i := 0 // 1. initial
		for {
			if i >= 10 { // 2. condition
				break
			}
			fmt.Println(i)
			i++ // 3. increment
		}
	}
}
