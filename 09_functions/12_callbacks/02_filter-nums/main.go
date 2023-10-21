package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var result []int // array
	for _, num := range numbers {
		if callback(num) {
			fmt.Println(num, "is division by 3.")
			result = append(result, num) // array
		}
	}
	return result
}

func main() {
	xs := filter([]int{99, 45, 17, 11, 2, 33, 8}, func(n int) bool {
		return n%3 == 0
	})

	fmt.Println(xs)
}
