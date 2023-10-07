package ref

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_timeMe(t *testing.T) {
	timed := makeTimedFunc(timeMe).(func(int) int)
	result := timed(2)
	assert.Equal(t, 4, result)
}
