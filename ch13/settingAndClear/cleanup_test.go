package settingAndClear

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "tempFile", fName)

	content, err := ioutil.ReadFile("testdata/hi.txt")
	if err != nil {
		fmt.Errorf("error %v", err)
	}
	assert.Equal(t, "hello", string(content))

}
