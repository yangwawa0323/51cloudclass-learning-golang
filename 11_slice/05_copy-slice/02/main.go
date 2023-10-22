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

	var copied []string = make([]string, len(greeting))
	copy(copied, greeting) // back origin slice up.
	copied[2] = "我好"
	fmt.Println(copied)
	fmt.Println(greeting)

}
