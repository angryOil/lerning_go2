package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getZeroMyError(t *testing.T) {
	zeroValueErr := getZeroMyError(true)
	assert.Nil(t, zeroValueErr)

	notZeroValueErr := getZeroMyError(false)
	assert.NotNil(t, notZeroValueErr)
}
