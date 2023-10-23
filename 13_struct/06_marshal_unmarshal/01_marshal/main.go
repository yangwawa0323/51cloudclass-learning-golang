// JSON
// {  "Name" : "yangwawa" , "Fridents" : [ "" , "" ]   }
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int // lowercase , private
}

func main() {
	p1 := person{"James", "Bouds", 20, 007}
	bs, _ := json.Marshal(p1) // golang struct --> json
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs)) // char 'A'  --> uint8  2^8  0-255 []uint8
}
