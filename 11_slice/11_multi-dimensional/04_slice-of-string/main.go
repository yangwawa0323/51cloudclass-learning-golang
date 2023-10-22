package main

import "fmt"

func main() {
	var records [][]string // len 0

	// student1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	// store the record
	// records = append(records, student1)

	// student2
	student2 := make([]string, 3)
	student2[0] = "Lisa"
	student2[1] = "92.00"
	student2[2] = "109.00"

	records = append(records, student2, student1)

	student3 := make([]int, 2)
	student3[0] = 100
	student3[1] = 22

	// records = append(records, student3) // 👎

	fmt.Println(records)
}
