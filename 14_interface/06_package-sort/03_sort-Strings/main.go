package main

import (
	"fmt"
	"sort"
u

type people []string // alias

func main() {
	studyGroup := people{"ZhangSan", "LiSi", "WangWu", "ZhaoLiu"}

	fmt.Println("Before sort: ", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("After sort: ", studyGroup)
}
