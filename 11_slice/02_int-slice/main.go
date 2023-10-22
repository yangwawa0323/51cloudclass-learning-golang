package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 9, 11, 13}

	for idx, val := range xs {
		fmt.Printf("%d - %d\n", idx, val)
	}
}
