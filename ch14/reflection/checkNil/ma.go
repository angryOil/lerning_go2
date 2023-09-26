package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	var m map[int]int
	fmt.Println(hasNilValue(m))
	fmt.Println(hasNilValue(i))
	fmt.Println(hasNilValue(3))
}

func hasNilValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		fmt.Println("novalue")
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		fmt.Println("one of kind")
		return iv.IsNil()
	default:
		fmt.Println("default")
		return false
	}
}
