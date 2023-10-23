package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour!",
		2: "你好",
		3: "Bongiorno!",
	}

	delete(myGreeting, 7) // 🦖 key 7 is not exists

	// if val, exists := myGreeting[2]; exists {
	// 	fmt.Println("That value exists")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("exists: ", exists)
	// } else {
	// 	fmt.Println("That value does not exists")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("exists: ", exists)
	// }

	fmt.Println(myGreeting)
}
