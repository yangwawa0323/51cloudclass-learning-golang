package main

import (
	"fmt"
	"sort"
)

type people []string // alias

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	studyGroup := people{"ZhangSan", "LiSi", "WangWu", "ZhaoLiu"}

	fmt.Println("Before sort: ", studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	// sort.Sort(studyGroup)
	fmt.Println("After sort: ", studyGroup)
}
