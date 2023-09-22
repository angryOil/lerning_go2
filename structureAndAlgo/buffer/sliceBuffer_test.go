package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})

	assert.Equal(t, 4, buf.ReadAble())
}

func TestRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.ReadAble())

	assert.Equal(t, 4, buf.ReadAble())
	readData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readData[i])
	}

	assert.Equal(t, 0, buf.ReadAble())
}

func TestWriteAndRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.ReadAble())

	buf.Write([]byte{5, 6})
	assert.Equal(t, 6, buf.ReadAble())

	readData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readData[i])
	}
	assert.Equal(t, 2, buf.ReadAble())

	buf.Write([]byte{7, 8, 9})
	assert.Equal(t, 5, buf.ReadAble())

	readData = buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+5), readData[i])
	}

	assert.Equal(t, 1, buf.ReadAble())

}
