package main

import "fmt"

func zero(val *int) { // val = 50    thread A  *val = 49  thread B *val = 49
	fmt.Println(val)
	fmt.Printf("%p\n", val)
	*val = 0
}

func main() {
	x := 50
	fmt.Println(&x)
	fmt.Printf("%p\n", &x)
	zero(&x)       // x = 50
	fmt.Println(x) // x = 0 ?
}
