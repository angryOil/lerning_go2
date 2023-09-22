package main

import "fmt"

func main() {
	var i interface{}
	var mine myInt = 20
	i = mine
	i2, ok := i.(string)
	if !ok {
		fmt.Printf("unexpected type for %v", i)
	}
	fmt.Println(i2)
}

type myInt int
