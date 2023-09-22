package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceStack_Pop(t *testing.T) {
	ss := NewSliceStack[int]()
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)

	assert.Equal(t, 3, ss.Pop())
	assert.Equal(t, 2, ss.Pop())
	assert.Equal(t, 1, ss.Pop())
}
