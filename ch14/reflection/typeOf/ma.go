package main

import (
	"fmt"
	"reflect"
)

func main() {
	fir()
	sec()
}

type Foo struct {
}

func fir() {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())

	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
	fmt.Println(xt.Kind())
	fmt.Println(ft.Kind())
	fmt.Println(xpt.Elem().Kind())
}

type per struct {
	Age  int
	name string
}

func sec() {
	var p per
	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		curField := pt.Field(i)
		fmt.Println(curField.Name, curField.Type.Name())
	}
}
