package main

import (
	"fmt"
	"reflect"
)

func main() {
	useFilter()
}

func useFilter() {
	names := []string{"hello", "joy", "world", "warm", "oil", "angryOil"}
	longNames := filter(names, func(s string) bool {
		return len(s) >= 4
	}).([]string)
	fmt.Print(longNames)

	ages := []int{1, 2, 45, 33, 22, 11, 30, 18, 16, 24}
	adult := filter(ages, func(a int) bool {
		return a >= 20
	})
	fmt.Println(adult)
}

func filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			fmt.Print(values[0].Kind() == reflect.Bool)
			fmt.Println(values[0].Type())
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}
