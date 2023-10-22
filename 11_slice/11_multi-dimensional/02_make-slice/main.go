package main

import "fmt"

func main() {
	var student []string = make([]string, 35)
	var students [][]string = make([][]string, 5)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // IMPORTANT!!!
	// A slice without initialized value has not memory allocation.
	fmt.Println(students == nil)
}
