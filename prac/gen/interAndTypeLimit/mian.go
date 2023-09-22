package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mi myIntStringer = 22
	fmt.Println(mi.String())
	print1(mi)
}

type intStringer interface {
	~int | ~int32 | ~int64
	String() string
}

type myIntStringer int

func (mi myIntStringer) String() string {
	return strconv.Itoa(int(mi))
}

func print1(a stringer) {
	fmt.Println(a.String())
}

func print2[T stringer](a T) {
	fmt.Println(a.String())
}

type stringer interface {
	String() string
}

type integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type myStr struct {
	name string
}

func (m myStr) String() string {
	return m.name
}
