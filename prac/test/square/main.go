package main

import "fmt"

func main() {
	fmt.Println("9 * 9 = ", square(9))
}

func square(i int) int {
	return i * i
}

func tdd(i int) int {
	return i * i
}
