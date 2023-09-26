package settingAndClear_test

import (
	"github.com/stretchr/testify/assert"
	"lerning_go2/ch13/settingAndClear"
	"testing"
)

func TestPublicStr(t *testing.T) {
	str := settingAndClear.PublicStr()
	assert.Equal(t, "public", str)
	// public 함수만 호출 가능
	//str = settingAndClear.privateStr()
	//assert.Equal(t, "private", str)
}
