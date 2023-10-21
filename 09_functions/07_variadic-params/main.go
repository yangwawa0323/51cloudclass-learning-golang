package main

import "fmt"

func main() {
	avg := average(1000, 100, 99, 45, 8)
	fmt.Printf("average: %.2f\n", avg)
}

// func average(start, input ...int) 👎
// func average(input ...int, start int) 👎 float32 { variadic parameter only be the last
func average(start int, input ...int) float32 { // 0 .. unlimit

	for _, num := range input {
		// total = total + num
		start += num
	}
	return float32(start) / float32(len(input))
}
