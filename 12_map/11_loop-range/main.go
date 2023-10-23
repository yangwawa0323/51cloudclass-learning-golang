package main

import "fmt"

func main() {
	myHost := map[string]string{
		"address":  "192.168.0.1",
		"username": "root",
		// "port" : 8080 // 👎
		"password": "secret",
		"database": "testing",
	}

	for key, val := range myHost {
		fmt.Println(key, " - ", val)
	}
}
