package main

import "fmt"

func main() {
	var i int
	getTypeName("str")
	getTypeName(i)
	getTypeName(1.2)
	getTypeName(nil)
}

func getTypeName(i interface{}) {
	switch j := i.(type) {
	case int:
		fmt.Println(i, " type is int ", j)
	default:
		fmt.Println(i, " type is ", j)
	}
}
