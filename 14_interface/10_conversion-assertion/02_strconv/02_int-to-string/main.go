package main

import (
	"fmt"
	"strconv"
)

func main() {
	//  "500hats" 👎  "hat500" 👎  "500.107" 👎
	var x = 500 // 💯
	y := "I have " + strconv.Itoa(x) + " hats"
	fmt.Println(y)

	var x2 = 600
	y2 := fmt.Sprintf("I have %d hats", x2)
	fmt.Println(y2)
}
