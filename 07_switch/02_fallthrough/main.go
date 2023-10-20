package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Go home")
		fallthrough
	case time.Sunday:
		fmt.Println("It's the weekend!!!")
	default:
		fmt.Println("It's a workday")
	}
}
