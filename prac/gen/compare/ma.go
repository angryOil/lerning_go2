package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

func main() {
	useCom()
	useMa()
}

func useMa() {
	doubled := ma([]int{1, 2, 3}, func(f int) int {
		return f * 2
	})
	fmt.Println(doubled)

	uppered := ma([]string{"hello", "world", "i`m", "joy"}, func(v string) string {
		return strings.ToUpper(v)
	})
	fmt.Println(uppered)

	toString := ma([]int{1, 2, 3, 4, 5}, func(i int) string {
		return "str" + strconv.Itoa(i)
	})
	fmt.Println(toString)
}

func ma[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

func useCom() {
	var m1 myStr = "hello"
	var m2 myStr = "world"
	var m3 myStr = "world"
	fmt.Println(equal(m1, m2))
	fmt.Println(equal(m2, m3))
}

type comparableHasher interface {
	comparable
	Hash() uint32
}

type myStr string

func (s myStr) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func equal[T comparableHasher](a, b T) bool {
	if a == b {
		return true
	}
	return a.Hash() == b.Hash()
}
