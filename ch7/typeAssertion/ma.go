package main

import (
	"fmt"
)

type myInt int

func main() {
	assertion1()

	assertion2(myInt(1))
	assertion2("hwllo")
}

func assertion2(i interface{}) {
	switch t := i.(type) {
	case nil:
		fmt.Println(t, "is nil")
	case myInt:
		fmt.Println(t, " is myint", i)
	default:
		fmt.Println("i don`t know ")
	}
}

func assertion1() {
	var i interface{}
	var mi myInt = 22
	i = mi
	i2 := i.(myInt)
	fmt.Println(i2)
	i3 := int(mi)
	i4 := string(mi)
	fmt.Println(i3)
	fmt.Println(i4)
	//i3 := i.(string)
	//fmt.Println(i3)
}
