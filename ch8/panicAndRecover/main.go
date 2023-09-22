package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{1, 2, 3, 0, 3, 4, 5}
	for _, n := range nums {
		div(n)
	}
}

func div(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Print(reflect.TypeOf(v))
			fmt.Println(v.(error).Error())
		}
	}()
	fmt.Println(60 / i)
}
