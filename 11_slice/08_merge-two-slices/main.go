package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{10, 100, 1000, 10000}

	// mySlice => [1,2,3,4,5, 10,100,1000,10000]
	// ...
	// mySlice = append(mySlice, myOtherSlice...) // append other slice after myself
	// fmt.Println(mySlice)

	// mySlice => [ 10, 100, 1000, 10000, 1,2,3,4,5]
	// mySlice = append(myOtherSlice... , mySlice) // 👎
	mySlice = append(myOtherSlice, mySlice...)
	fmt.Println(mySlice)
}
