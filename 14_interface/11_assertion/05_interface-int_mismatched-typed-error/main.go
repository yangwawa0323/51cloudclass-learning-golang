package main

import "fmt"

func main() {
	var name interface{} = "Yangwawa"
	var val interface{} = 7
	fmt.Printf("val: %T , name: %T\n", val, name)
	fmt.Println(val + 6)
}
