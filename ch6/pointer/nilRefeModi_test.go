package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_failToUpdate(t *testing.T) {
	f := 1
	failToUpdate(&f)
	assert.Equal(t, 10, f)
}

func Test_successToUpdate(t *testing.T) {
	s := 1
	successToUpdate(&s)
	fmt.Println("::", s)
	assert.Equal(t, 10, s)
}

type per struct {
	Name string
}
