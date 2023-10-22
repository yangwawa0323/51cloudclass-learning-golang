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

	fmt.Println("greeting[1:3]", greeting[1:3])
	fmt.Println("greeting[:6]", greeting[:6]) // remaining
}
