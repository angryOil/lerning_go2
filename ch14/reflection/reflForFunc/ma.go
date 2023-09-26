package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	wrappedIntFunc := printDuringTime(toDoubleScore).(func(int2 int) int)
	fmt.Print(wrappedIntFunc(2))
	wrappedStringFunc := printDuringTime(justReturnOriginStr).(func(string2 string) string)
	fmt.Print(wrappedStringFunc("wow"))
}

func justReturnOriginStr(origin string) string {
	return origin
}

func toDoubleScore(x int) int {
	return x * x
}

func printDuringTime(f interface{}) interface{} {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		fmt.Println("in?", in)
		out := fv.Call(in)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}
