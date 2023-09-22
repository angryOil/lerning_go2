package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Pop())
	assert.Equal(t, 4, q.Pop())
}

func BenchmarkNewQueue(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
func BenchmarkNewSliceQueue(b *testing.B) {
	s := NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
