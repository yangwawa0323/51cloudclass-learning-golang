package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name   string
	Age    int
	Height float32
	Weight float32
}

type ByAge []person // alias

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// ///////////////////////////
type ByHeight []person // alias

func (h ByHeight) Len() int { // implemented sort.Interface
	return len(h)
}

func (h ByHeight) Less(i, j int) bool {
	return h[i].Height < h[j].Height
}

func (h ByHeight) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	people := []person{
		{Name: "Bob", Age: 31, Height: 188},
		{Name: "John", Age: 24, Height: 172.5},
		{Name: "Michael", Age: 42, Height: 182.0},
		{Name: "Jenny", Age: 26, Height: 168.4},
	}
	fmt.Println("Before sort by age: ", people)

	sort.Sort(ByAge(people))

	fmt.Println("After sort by age: ", people)

	sort.Sort(sort.Reverse(ByHeight(people)))

	fmt.Println("After sort by height: ", people)
}
