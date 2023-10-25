package main

import "fmt"

type Door struct {
	Color  string
	Height float32
	width  float32
}

func (d Door) open() {
	fmt.Println("I am door")
}

func (d Door) close() {
	fmt.Println("closing door")
}

func (d Door) watch() {

}

type Closeable interface {
	open()
	close()
	watch()
}

func into(c Closeable) {
	c.open()
	c.close()
}

func main() {
	d := Door{
		"Yellow", 220.00, 150.00,
	}
	into(d)
}
