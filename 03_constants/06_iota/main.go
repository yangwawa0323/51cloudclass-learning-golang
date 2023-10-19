package main

import "fmt"

const (
	HTTP_NOT_SEARCHABLE = iota + 401
	HTTP_NOT_REACHABLE  // 402
	HTTP_NOT_SUCH_FILE  // 403
	HTTP_NOT_FOUND      // 404
)

func main() {
	fmt.Println("HTTP_NOT_FOUND", HTTP_NOT_FOUND)
	fmt.Println("HTTP_NOT_SUCH_FILE", HTTP_NOT_SUCH_FILE)
}
