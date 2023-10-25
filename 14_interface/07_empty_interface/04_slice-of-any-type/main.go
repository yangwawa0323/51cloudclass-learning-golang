package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"Woof"}, true}
	fifi := cat{animal{"Meow"}, true}
	shadow := dog{animal{"Woof"}, true}
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)

}
