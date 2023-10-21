package main

import "fmt"

func main() {
	var x [256]int // 0-255

	fmt.Println(len(x))
	for i := 0; i < 256; i++ {
		x[i] = i // x[0] = 0 , x[1] = 1 .. x[255] = 255
	}

	for _, v := range x {
		fmt.Printf("%v - %T - %08b\n", v, v, v)
		// %08b    0 - zero-fill, 8 - min length , b - binary
	}

	// fmt.Println(x)
}
