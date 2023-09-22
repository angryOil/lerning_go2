package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRingWrite(t *testing.T) {
	buf := NewRingBuffer[byte](1024)

	writed := buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, writed)
	assert.Equal(t, 4, buf.ReadAble())
}

func TestRingRead(t *testing.T) {
	buf := NewRingBuffer[byte](1024)
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.ReadAble())

	assert.Equal(t, 4, buf.ReadAble())
	readData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readData[i])
	}

	assert.Equal(t, 0, buf.ReadAble())
}

func TestRingOverWrite(t *testing.T) {
	buf := NewRingBuffer[byte](5)
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.ReadAble())
	assert.Equal(t, 4, buf.writePoint)
	assert.Equal(t, 0, buf.readPoint)
	assert.Equal(t, 1, buf.Writeable())

	buf.Write([]byte{5})
	assert.Equal(t, 5, buf.ReadAble())
	assert.Equal(t, 0, buf.writePoint)
	assert.Equal(t, 0, buf.Writeable())

	writed := buf.Write([]byte{6})
	assert.Equal(t, 5, buf.ReadAble())
	assert.Equal(t, 0, writed)
	assert.Equal(t, 0, buf.Writeable())

	readData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readData[i])
	}
	assert.Equal(t, 1, buf.ReadAble())
	assert.Equal(t, 4, buf.Writeable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(t, 3, writed)
	assert.Equal(t, 3, buf.writePoint)
	assert.Equal(t, 4, buf.ReadAble())
	assert.Equal(t, 1, buf.Writeable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(t, 1, writed)
	//assert.Equal(t, 4, buf.writePoint)
	assert.Equal(t, 5, buf.ReadAble())
	assert.Equal(t, 0, buf.Writeable())

	readData = buf.Read(4)
	assert.Equal(t, 4, len(readData))
	assert.Equal(t, 1, buf.ReadAble())
	assert.Equal(t, 4, buf.Writeable())
}

//func TestRingWriteAndRead(t *testing.T) {
//	buf := NewSliceBuffer[byte]()
//	buf.Write([]byte{1, 2, 3, 4})
//	assert.Equal(t, 4, buf.ReadAble())
//
//	buf.Write([]byte{5, 6})
//	assert.Equal(t, 6, buf.ReadAble())
//
//	readData := buf.Read(4)
//	for i := 0; i < 4; i++ {
//		assert.Equal(t, byte(i+1), readData[i])
//	}
//	assert.Equal(t, 2, buf.ReadAble())
//
//	buf.Write([]byte{7, 8, 9})
//	assert.Equal(t, 5, buf.ReadAble())
//
//	readData = buf.Read(4)
//	for i := 0; i < 4; i++ {
//		assert.Equal(t, byte(i+5), readData[i])
//	}
//
//	assert.Equal(t, 1, buf.ReadAble())
//
//}
