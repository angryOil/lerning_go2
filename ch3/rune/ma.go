package main

import (
	"fmt"
	"strconv"
)

func main() {
	ruteTest()
}

type per struct {
	name string
}

func ruteTest() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	var num = 112
	var s3 = strconv.Itoa(num)

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(s2)
	fmt.Println(s3)

	var p = per{
		name: "wow",
	}
	var anotherStruct = struct {
		name string
	}{
		name: "wow",
	}
	fmt.Println(p == anotherStruct)
	var ap per = anotherStruct
	fmt.Println(p == ap)
}
