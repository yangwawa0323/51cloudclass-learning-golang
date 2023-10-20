package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var url string = "https://www.bing.com/notexists.html"

	response, _ := http.Get(url)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
