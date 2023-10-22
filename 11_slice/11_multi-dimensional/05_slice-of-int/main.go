package main

import "fmt"

func main() {
	queues := make([][]int, 0, 3) // len 0 , cap 3

	for i := 0; i < 3; i++ {
		row := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			row = append(row, j)
		}
		queues = append(queues, row)
	}

	fmt.Println(queues)
}
