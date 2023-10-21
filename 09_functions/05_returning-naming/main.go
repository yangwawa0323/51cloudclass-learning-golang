package main

import "fmt"

func main() {
	fmt.Println(greet("Yang", "wawa"))
}

func greet(fname, lname string) (s string) { // ""
	s = fmt.Sprintf("%s %s , nice to meet you.", fname, lname)
	// var _ = fullname
	return
}
