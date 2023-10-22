package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	myArray := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("mySlice %T\n", mySlice)
	fmt.Printf("myArray %T\n", myArray)
	fmt.Println(mySlice)
	fmt.Println(myArray)
}
