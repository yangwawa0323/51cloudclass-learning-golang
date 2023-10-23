package main

import "fmt"

// struct fields , miss field use default , age int default is 0
type person struct {
	first string
	last  string
	age   int
}

// reciever
func (p person) fullname() string {
	return fmt.Sprintf("%s , %s is %d age.", p.first, p.last, p.age)
}

func main() {
	p2 := person{
		last:  "Moneypenny",
		first: "Miss",
	}

	fmt.Println(p2.fullname())
}
