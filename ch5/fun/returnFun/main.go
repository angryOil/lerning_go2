package main

import (
	"fmt"
	"os"
)

func main() {
	sevenBase := makeMulti(7)
	fmt.Println(sevenBase(22))
	readFile()
}

func makeMulti(base int) func(int) int {
	return func(i int) int {
		return base * i
	}
}

func readFile() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	ar := os.Args
	fmt.Println(ar)
}
