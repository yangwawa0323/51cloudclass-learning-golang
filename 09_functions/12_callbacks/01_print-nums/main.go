// anonymous

package main

import "fmt"

type callbackFunc func(int)

func visit(numbers []int, callback callbackFunc) {
	/// code
	for _, num := range numbers {
		callback(num)
	}
}

func main() {
	// visit( ) 1. array 2. anonymous
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		if n%2 == 0 {
			fmt.Println(n)
		}
	})
}
