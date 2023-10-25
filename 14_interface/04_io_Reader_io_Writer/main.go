package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type FakeReaderWriter struct{}

func (f FakeReaderWriter) Read(p []byte) (n int, err error) {
	// Logical
	return 109, nil
}

func (f FakeReaderWriter) Write(p []byte) (n int, err error) {
	// Write database
	return 109, nil
}

func main() {
	buffer := bytes.NewReader([]byte("Reader Writer io.Copy example\n\n"))

	var bioReader = bufio.NewReader(buffer) // many time with seek
	var bioWriter = bufio.NewWriter(os.Stdout)
	written, err := io.Copy(bioWriter, bioReader)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("written to os.Stdout: ", written)

	/////////////////////
	fk := FakeReaderWriter{}
	written, err = io.Copy(fk, bioReader)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("written to FakeReaderWriter: ", written)
}
