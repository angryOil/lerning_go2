package main

import "fmt"

func main() {
	a()
}

func a() {
	var c = make([]int32, 0, 0)
	fmt.Println(cap(c))
	fmt.Println(len(c))
	c = append(c, 11)
	fmt.Println(cap(c))
	fmt.Println(len(c))
}
