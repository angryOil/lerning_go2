package ch13

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_TestData(t *testing.T) {
	data, err := os.ReadFile("testdata/test.file")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "이것은 테스트 파일입니다.", string(data))
}
