// yangwawa ["Linux"]
// Jake ["Linux", "MySQL" , "Oracle"]   []string
package main

import "fmt"

func main() {
	var published = make(map[string][]string)

	published["yangwawa"] = []string{"Linux"}
	published["jake"] = make([]string, 0)
	// fmt.Printf("%T published[jake]\n", published["jake"])
	published["jake"] = append(published["jake"], "Beginning of Linux")
	published["jake"] = append(published["jake"], "MySQL administrate")

	fmt.Println(published)
}
