package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 * 10 -->  1000000000
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println("Binary\t\tDecimal")
	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)
}
