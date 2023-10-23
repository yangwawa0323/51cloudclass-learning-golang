package main

import "fmt"

// SQL
//  CREATE TABLE people (
// first  varchar(255),
// last varchar(255),
// age  int unsign
//)
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20} // OK
	// p1 := person{"Bond", "James", 20} // not OK
	// p1 := person{20, "Bond", "James"} // not OK

	p2 := person{
		age:   18,
		last:  "Moneypenny",
		first: "Miss",
	}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
