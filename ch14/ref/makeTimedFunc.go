package ref

import (
	"fmt"
	"reflect"
	"time"
)

func timeMe(a int) int {
	time.Sleep(1 * time.Second * time.Duration(a))
	result := a * 2
	return result
}

func makeTimedFunc(f interface{}) interface{} {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(in)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}
