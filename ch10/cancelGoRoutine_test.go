package ch10

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countTo(t *testing.T) {
	ch, cancel := countTo(10)
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		result := <-ch
		fmt.Println("i:", i, "\tresult", result)
		assert.Equal(t, i, result)
	}
	cancel()
}
