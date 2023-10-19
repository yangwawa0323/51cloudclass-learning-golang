package main

import "fmt"

const metersToYards float64 = 1.1093

func main() {
	var meters float64 // default : 0
	fmt.Print("PLS Enter meters: ")
	fmt.Scan(&meters) //
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
