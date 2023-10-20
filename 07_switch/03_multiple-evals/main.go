package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!!!")
	default:
		fmt.Println("It's a workday")
	}

	var name = "Jake"
	switch name {
	case "Daniel", "Yangwawa":
		fmt.Println("It is long time no see", name)
	case "Jenny", "Jake":
		fmt.Println("Hi, it is other school")
	default:
		fmt.Println("Does we meet before?")
	}
}
