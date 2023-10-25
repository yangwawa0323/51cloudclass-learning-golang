package main

import "fmt"

type person struct {
	Name   string
	Age    int
	Weight float32
}

func (p person) LostWeight(w float32) {
	fmt.Println("Before exercism: ", p.Weight)
	p.Weight = p.Weight - w
	fmt.Println("After exercism: ", p.Weight)
}

type dog struct {
	Name     string
	Weight   float32
	Friendly bool
}

type life interface {
	LostWeight(w float32)
}

func exercism(l life, w float32) {
	l.LostWeight(w)
}

func main() {
	yk := person{Name: "Yangwawa", Age: 49, Weight: 75.6}

	exercism(yk, 4.4)
	fmt.Println(yk)
}
