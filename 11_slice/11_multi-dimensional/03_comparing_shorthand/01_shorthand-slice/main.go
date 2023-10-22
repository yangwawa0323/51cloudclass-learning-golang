package main

import "fmt"

func main() {
	student := []string{} // len = 0
	students := [][]string{}

	fmt.Println("student len: ", len(student))
	fmt.Println("students len: ", len(students))
	student[0] = "Todd" // 👎  runtime error  index out of range

	fmt.Println(student)
	fmt.Println(students)
}
