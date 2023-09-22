package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_square(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be  81 but return %d", rst)
	}
}

func Test_square2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should be  9 but return %d", rst)
	}
}

func Test_square3(t *testing.T) {
	rst := square(0)
	if rst != 0 {
		t.Errorf("square(0) should be  0 but return %d", rst)
	}
}

func Test_withAssertion(t *testing.T) {
	assert.Equal(t, 81, square(9), "square(9) should be 81 but returns %d", square(9))
}

func Test_tdd(t *testing.T) {
	assert.Equal(t, 9, tdd(3), "should be 9")
	assert.Equal(t, 81, tdd(9), "should be 81")
	assert.Equal(t, 16, tdd(4), "should be 16")
	assert.Equal(t, 1, tdd(-1), "should be 1")
}
