package withTimeOut

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_doPrentCtxCancelAndChildCtxCancel(t *testing.T) {
	timeStr := doPrentCtxCancelAndChildCtxCancel()
	fmt.Println(timeStr)
	firstTime := []rune(timeStr)[0]
	assert.Equal(t, "2", string(firstTime))
}
