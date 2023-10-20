package main

import "fmt"

func main() {
	// if num := 19; num < 0 {
	// 	fmt.Println(num, "is negative.")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	myFriendsName := "Jake"
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("my friend name length is 2")
	case myFriendsName == "Tim":
		fmt.Println("Nice to meet you, Tim")
	case myFriendsName == "Marcus", myFriendsName == "John":
		fmt.Println("It is long time no see", myFriendsName)
	default:
		fmt.Println("It is stranger")
	}
}
