package main

import "fmt"

func main() {
	// mySlice := make([]int, 0, 5) // `0` -> initial length, `3` -> capacity
	mySlice := make([]int, 3) // `3` -> initial length, `3` capacity
	fmt.Println("================")
	fmt.Println(mySlice)
	fmt.Println("len: ", len(mySlice))
	fmt.Println("cap: ", cap(mySlice))
	fmt.Println("================")

	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice),
			"value: ", mySlice[i])
	}
}
