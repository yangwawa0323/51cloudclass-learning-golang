package main

import "fmt"

func main() {
	var students = []string{"ZhangSan", "LiSi", "WangWu"}

	for _, std := range students {

		func(name string) {
			// defined variable
			fmt.Println("Hi! ", name)
		}(std)
	}
}
