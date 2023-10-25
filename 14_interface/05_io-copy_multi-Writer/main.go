package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type FakeReaderWriter struct{} // is a Reader/Writer

func (f FakeReaderWriter) Read(p []byte) (n int, err error) {
	// Logical
	return len(p), nil
}

func (f FakeReaderWriter) Write(p []byte) (n int, err error) {
	// Write database
	fmt.Println("Writing message to faker now...", string(p))
	return len(p), nil
}

func main() {
	buffer := bytes.NewReader([]byte("Reader Writer io.Copy example\n\n"))

	file, err := os.Create("io-copy.txt")

	if err != nil {
		log.Panic(err)
	} else {
		defer file.Close()
	}

	////////
	fk := FakeReaderWriter{}

	written, err := io.Copy(io.MultiWriter(fk, file), buffer)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("written to :", written)

}
