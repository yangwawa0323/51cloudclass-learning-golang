package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string `json:"first,omitempty"`
	Last        string `json:"last,omitempty"`
	Gender      string `json:"sex,omitempty"`
	Age         int    `json:"age,omitempty"`
	notExported int    `json:"not_exported,omitempty"` // lowercase , private, visibility
}

func main() {
	var p1 person // ?
	fmt.Println(p1)

	fmt.Println(p1.First)

	// data is coming from server side
	// json bytes convert into struct
	data := []byte(`{"last":"Bouds","sex": "女", "age":20,"NotExported":7}`)
	if err := json.Unmarshal(data, &p1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)

	// struct convert into json bytes
	p2 := person{
		First:  "MoneyPenny",
		Last:   "Miss",
		Age:    20,
		Gender: "男",
	}
	bs, _ := json.Marshal(p2)
	fmt.Println(string(bs))
}
