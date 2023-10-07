package ch13

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "tempFile", fName)
}

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		fmt.Println("삭제를 시도합니다")
		os.Remove(f.Name())
	})
	return f.Name(), nil
}
