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
		"Gutten morgen!"}

	fmt.Println(greeting)
	greeting = append(greeting[:2], greeting[3:]...)

	fmt.Println(greeting)

}
