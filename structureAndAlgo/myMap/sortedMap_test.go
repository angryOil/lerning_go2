package myMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedMap_Get(t *testing.T) {
	var s sortedMap[string, int]

	s.Add("joy", 923)
	v, ok := s.Get("joy")
	assert.True(t, ok)
	assert.Equal(t, 923, v)

	s.Add("world", 22)
	v, ok = s.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 22, v)
	//s.Add("joy", 3)
	//s.Add("joy", 3)
	//s.Add("joy", 3)
	//assert.Equal(t, 3, s.Get("joy"))
	//assert.Equal(t, 3, s.Get("joy"))

	assert.Equal(t, "joy", s.Arr[0].key)
	assert.Equal(t, "world", s.Arr[1].key)
	//assert.Equal(t, "world", s.Arr[2].key)
}

func TestOverWarp(t *testing.T) {
	var s sortedMap[string, int]
	s.Add("bbb", 10)
	v, ok := s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 10, v)

	s.Add("bbb", 20)
	v, ok = s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 20, v)
	assert.Equal(t, 1, len(s.Arr))

}

func TestGetEmpty(t *testing.T) {
	var s sortedMap[string, int]
	s.Add("bbb", 10)
	v, ok := s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 10, v)

	v, ok = s.Get("aaaa")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestRemove(t *testing.T) {
	var s sortedMap[string, int]
	s.Add("ccc", 30)
	s.Add("bbb", 20)
	s.Add("aaa", 10)

	removed := s.Remove("bbb")
	assert.True(t, removed)

	removed = s.Remove("bbb")
	assert.False(t, removed)

	assert.Equal(t, 2, len(s.Arr))

	assert.Equal(t, "aaa", s.Arr[0].key)
	assert.Equal(t, "ccc", s.Arr[1].key)

}
