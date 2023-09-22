package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	assert.Equal(t, 4, s.Pop())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())
}

func BenchmarkNewStack(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
func BenchmarkNewSliceStack(b *testing.B) {
	s := NewSliceStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
