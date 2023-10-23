package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int // lowercase , private, visibility
}

func main() {
	var p1 person // ?
	fmt.Println(p1)
	fmt.Println(p1.First)

	// data is coming from server side
	data := []byte(`{"First":"James","Last":"Bouds","Age":20,"NotExported":7}`)
	json.Unmarshal(data, &p1)
	fmt.Println(p1)
}
