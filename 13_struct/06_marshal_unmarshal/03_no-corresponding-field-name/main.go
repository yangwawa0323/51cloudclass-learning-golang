package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Gender      uint8 // unsigned   0: female 1: male
	Age         int
	notExported int // lowercase , private, visibility
}

func main() {
	var p1 person // ?
	fmt.Println(p1)
	fmt.Println(p1.First)

	// data is coming from server side
	data := []byte(`{"first":"James","last":"Bouds","sex",1 "Age":20,"NotExported":7}`)
	json.Unmarshal(data, &p1)
	fmt.Println(p1)
}
