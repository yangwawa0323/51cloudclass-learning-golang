package main

import "fmt"

func main() {
	x := 42
	fmt.Println("main : ", x)
	{ // closure  lifetime
		x := 90
		fmt.Println("closure: ", x)
		y := "This var is inside closure"
		fmt.Println(y)
	}
	// fmt.Println("main : ", y) 👎
	fmt.Println("main : ", x)
}
