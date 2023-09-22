package main

import "fmt"

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	first := 1
	sec := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = first + sec
		sec = first
		first = rst
	}
	return rst
}
