package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{77, 8, 3, 4, 691, 498, 100, 0}

	fmt.Println("Before sort : ", n)
	// sort.Sort(sort.Reverse(sort.IntSlice(n))) // 👍
	sort.Ints(n) // 👎 cannot reverse
	fmt.Println("After sort : ", n)
}
