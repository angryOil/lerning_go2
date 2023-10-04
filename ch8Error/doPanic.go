package ch8Error

import "fmt"

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
