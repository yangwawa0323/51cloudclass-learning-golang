package main

import "fmt"

type Door struct {
	Color  string
	Height float32
	width  float32
}

func (d Door) open() {
	fmt.Println("open the door")
}

func (d Door) close() {
	fmt.Println("close the door")
}

func (d Door) authenticate(s string) bool {
	fmt.Println("your username", s)
	return true
}

type Machanical interface {
	open()
	close()
}

type Equipment interface {
	authenticate(secret string) bool
}

func into(m Machanical) {
	m.open()
	m.close()
}

func checkIn(name string, e Equipment) {
	if e.authenticate(name) {
		fmt.Println("Your are welcome")
	} else {
		fmt.Println("Access denied")
	}
}

func main() {
	d := Door{Color: "Yellow"}
	into(d)
	checkIn("yangwawa", d)
}
