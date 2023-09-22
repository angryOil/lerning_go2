package main

import "fmt"

func main() {
	runeFir()
	toSlice()
	loopToRuneSlice()
}

var s = "abc안녕하세요?!!"

func runeFir() {
	fmt.Println(s[1], s[4])
	fmt.Println(s[0:1])
	fmt.Println(string(s[0]))
	fmt.Println(string(s[1]))
	fmt.Println(string(s[1]))
}

func toSlice() {
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}

func loopToRuneSlice() {
	for _, r := range []rune(s) {
		fmt.Print(string(r))
	}
}
