package main

import "fmt"

func main() {
	var student []string // len 0
	var students [][]string

	fmt.Println("student len: ", len(student))
	fmt.Println("students len: ", len(students))
	student[0] = "Todd" // 👎  runtime error  index out of range

	fmt.Println(student)
	fmt.Println(students)
}
