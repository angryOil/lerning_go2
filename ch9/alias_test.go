package ch9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_foo_hello(t *testing.T) {
	var b = bar{x: 20}
	var hello = b.hello()
	assert.Equal(t, "20hello I`m foo", hello)
}
