package main

import (
	"fmt"
	"math"
)

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type rectangle struct {
	height float32
	width  float32
}

func (r rectangle) area() float32 {
	return r.height * r.width
}

type shape interface { // alias type
	// Any struct implemented all of functions in this
	//interface , must be this type.
	area() float32
}

func info(sp shape) { // 1. square 2. circle
	fmt.Println("shape area : ", sp.area())
}

func main() {
	s := square{side: 10}
	fmt.Printf("%T\n", s)
	c := circle{radius: 3.5}
	r := rectangle{height: 40, width: 20}
	fmt.Println(s.area())
	info(s) // s main.square struct
	info(c) // c main.circle struct
	info(r)
}
