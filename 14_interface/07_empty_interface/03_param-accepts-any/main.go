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

func specs(a any) { // a interface{}
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"Woof"}, true}
	fifi := cat{animal{"Meow"}, true}
	specs(fido)
	specs(fifi)
}
