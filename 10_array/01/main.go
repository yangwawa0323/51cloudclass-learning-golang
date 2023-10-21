package main

import "fmt"

func main() {
	var x [58]int // array fixed length
	// var y []int  // slice unknown length
	fmt.Println(x)
	// fmt.Println(y)
	fmt.Println(" x length : ", len(x))
	fmt.Println(x[42]) // 43th element  0-index
	x[2] = 999
	fmt.Println(x)
}
