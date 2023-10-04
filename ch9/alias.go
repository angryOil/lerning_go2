package ch9

import "strconv"

type foo struct {
	x int
	S string
}

func (f foo) hello() string {
	return strconv.Itoa(f.x) + "hello I`m foo"
}

type bar = foo
