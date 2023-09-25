package myMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hashMap_Add(t *testing.T) {
	var h hashMap[int]
	h.Add("joy", 30)

	val, ok := h.Get("joy")
	assert.True(t, ok)
	assert.Equal(t, 30, val)

	h.Add("world", 200)
	val, ok = h.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 200, val)

	val, ok = h.Get("joy")
	assert.True(t, ok)
	assert.Equal(t, 30, val)

	h.Add("goDang", 311)
	val, ok = h.Get("goDang")
	assert.True(t, ok)
	assert.Equal(t, 311, val)
}

func TestGoBasicMap(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 100
	m["b"] = 200
	m["c"] = 300
	m["d"] = 0

	assert.Equal(t, 100, m["a"])
	assert.Equal(t, 200, m["b"])
	assert.Equal(t, 300, m["c"])
	assert.Equal(t, 0, m["d"])

	noValue, ok := m["noValue"]
	assert.Equal(t, 0, noValue)
	assert.Equal(t, false, ok)
}
