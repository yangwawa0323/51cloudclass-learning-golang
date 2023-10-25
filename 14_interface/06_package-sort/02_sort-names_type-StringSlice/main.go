package main

import (
	"fmt"
	"sort"
)

type people []string // alias

func main() {
	studyGroup := people{"ZhangSan", "LiSi", "WangWu", "ZhaoLiu"}

	fmt.Println("Before sort: ", studyGroup)
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println("After sort: ", studyGroup)
}
