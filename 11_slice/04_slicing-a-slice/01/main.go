package main

import "fmt"

func main() {
	var results []int
	fmt.Println("results: ", results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}

	fmt.Println("mySlice: ", mySlice)

	fmt.Println("mySlice[2:4]", mySlice[2:4]) // 👍
	fmt.Println("mySlice[2:3]", mySlice[2:3]) // 👎
	fmt.Println("mySlice[2]", mySlice[2])

}
