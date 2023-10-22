package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!", //0
		"Bonjour!",      //1
		"你好",            //2
		"dias!",         // 3
		"Bongiorno!",    //4
		"Ohayo!",        //5
		"Selamat pagi!", // 6
		"Gutten morgen!",
	}

	copy := greeting[:] // In python copy, greeting difference memory address
	copy[2] = "我好"
	fmt.Println(copy)
	fmt.Println(greeting)

}
