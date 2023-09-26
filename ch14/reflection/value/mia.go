package main

import (
	"fmt"
	"reflect"
)

func main() {
	valueof()
	noValueNew()
	reflNewAndMakeSlice()
}

var stringSliceType = reflect.TypeOf([]string(nil))
var stringType = reflect.TypeOf((*string)(nil)).Elem()

func reflNewAndMakeSlice() {
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	sv := reflect.New(stringType).Elem()
	sv.SetString("new Str")

	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)
}

func noValueNew() {
	strType := reflect.TypeOf((*string)(nil)).Elem()
	fmt.Println(":", strType)
	fmt.Println(":", reflect.ValueOf(strType))
	sliceType := reflect.TypeOf([]string(nil))
	fmt.Println(sliceType)
	sv := reflect.New(strType).Elem()
	sv.SetString("hello")
	fmt.Println(sv.Type())
	fmt.Println(sv)
}

func valueof() {
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	s2 := sv.Interface().([]string)

	fmt.Println(sv.Type())
	fmt.Println(sv.Kind())
	fmt.Println(sv)
	fmt.Println(s2)

	i := 10
	iv := reflect.ValueOf(&i)
	fmt.Println(iv)
	fmt.Println(iv.Elem().Kind())
	fmt.Println(iv.Elem().Type())
	i2 := 20
	i2v := reflect.ValueOf(&i2)
	i2v.Elem().SetInt(3000)
	fmt.Println(i2)
}
