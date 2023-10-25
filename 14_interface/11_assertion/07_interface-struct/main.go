package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type dog struct {
	Name  string
	Sound string
}

func main() {
	var p interface{} = person{"Yangwawa", 49}
	var d interface{} = dog{"Moffy", "Woof"}

	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%s\n", p.(person).Name)
	fmt.Printf("%s\n", d.(dog).Name)
}
