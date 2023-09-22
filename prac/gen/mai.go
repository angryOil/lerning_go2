package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	useCustom()
	useCon()
}

func useCon() {
	var a = 3
	var b = 5
	fmt.Println(useConstant(a, b))
	var c = "hello"
	var d = "world"
	fmt.Println(useConstant(c, d))
}

func useCustom() {
	var a float32 = 22.13
	var b float32 = 22.53
	fmt.Println(min(a, b))

	var c int = 3
	var d int = 44
	fmt.Println(min(c, d))
}

func useAny() {
	var a int = 10
	print(a)
	var b float32 = 3.154
	var c string = "hello"
	print(b)
	print(c)
}

func print[T any](a T) {
	fmt.Println(a)
}

type float interface {
	float32 | float64
}

type myInt int8

// ~ << 이걸 포함하면 별칭도 포함 위에 myInt 같은경우도 포함 가능
type withAliasInteger interface {
	~int | ~int32 | ~int8 | ~uint16
}

type integer interface {
	int | int64 | int32 | int8
}

type num interface {
	integer | float
}

func useConstant[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func min[T num](a, b T) T {
	if a < b {
		return a
	}
	return b
}
