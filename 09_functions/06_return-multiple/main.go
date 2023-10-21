package main

import "fmt"

func main() {
	en, cn := greet("Yang", "wawa")
	fmt.Println(en)
	fmt.Println(cn)
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprintf("Hello! %s %s", fname, lname),
		fmt.Sprintf("你好! %s %s", fname, lname)
}
