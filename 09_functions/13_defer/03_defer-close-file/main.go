package main

import (
	"fmt"
	"os"
)

func main() {
	// open
	file, err := os.Open("09_functions/13_defer/03_defer-close-file/main.go")
	// error check
	if err != nil {
		panic(err)
	}
	// defer closing something. eg. file, socker, database.
	defer func() {
		fmt.Println("Closing file.")
		file.Close()
	}()
	///

	var buffer []byte = make([]byte, 1024)
	if _, err := file.Read(buffer); err != nil {
		fmt.Println("file read occur error :", err)
	}

	fmt.Println(string(buffer))

}
