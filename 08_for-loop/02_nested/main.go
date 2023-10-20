package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := i; j <= 9; j++ {
			fmt.Printf("%d * %d = %d \n", i, j, i*j)
		}
	}
}
