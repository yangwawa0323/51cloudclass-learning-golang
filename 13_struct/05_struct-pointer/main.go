package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeName(p *person) {
	p.name = "xxxx"
}

func main() {
	p1 := person{"James", 20}
	p2 := &person{"yangwawa", 49}
	fmt.Println(p1, p2)
	fmt.Printf("%T %T\n", p1, p2)

	changeName(&p1)
	changeName(p2)

	// p2.name = "YangKun"
	fmt.Println(p1, p2)
}
