package main

import "fmt"

func main() {

ExitSearchDir:
	for j := 0; j < 3; j++ { // find directoies from root dir
		fmt.Println("Search each directory")
		i := 0 // 1. initial
		// ExitSearchFile:
		for { // find file
			fmt.Println("Search each file")
			if i >= 10 { // 2. condition
				break ExitSearchDir
			}
			fmt.Println(i)
			i++ // 3. increment
		}
	}
}
