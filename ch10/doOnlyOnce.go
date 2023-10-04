package ch10

import (
	"fmt"
	"sync"
)

type slowComplicatedParser interface {
	Parse(string) string
}

var parser slowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
	once.Do(func() {
		fmt.Println("parser init!")
	})
	return ""
}
