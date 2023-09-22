package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceStack_Pop(t *testing.T) {
	sq := NewSliceQueue[int]()
	sq.Push(1)
	sq.Push(2)
	sq.Push(3)

	assert.Equal(t, 1, sq.Pop())
	assert.Equal(t, 2, sq.Pop())
	assert.Equal(t, 3, sq.Pop())
}
