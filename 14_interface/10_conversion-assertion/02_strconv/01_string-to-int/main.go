package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//  "500hats" 👎  "hat500" 👎  "500.107" 👎
	var x = "500" // 💯
	var y = 6
	z, err := strconv.Atoi(x)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(z + y)
}
