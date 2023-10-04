package ch8Error

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_doPanic(t *testing.T) {
	assert.Panics(t, func() { doPanic("panic!!!") }, "panic")
	assert.Panics(t, func() { fmt.Print("hello") }, "pnaic?")
}

func Test_div60(t *testing.T) {
	for _, val := range []int{1, 3, 5, 0, 4, 2} {
		if val == 0 {
			assert.Panics(t, func() {
				div60(val)
			}, "will be panic")
		}
		div60(val)
	}

}
