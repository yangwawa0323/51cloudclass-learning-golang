package main

import "fmt"

func main() {
	student := make([]string, 35) // ONLY use make function initialized the slice
	students := make([][]string, 35)

	fmt.Println("student len: ", len(student))
	fmt.Println("students len: ", len(students))
	student[0] = "Todd" //   runtime error  index out of range

	fmt.Println(student)
	fmt.Println(students)
}
