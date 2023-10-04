package ch10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_doDeadLock(t *testing.T) {
	assert.Panics(t, func() {
		doDeadLock()
	}, "dead rock!")
}

func Test_useSelect(t *testing.T) {
	useSelect()
}
