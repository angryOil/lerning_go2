package main

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {
	useSizeOf()
}

func useSizeOf() {
	var a int32 = 2
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(33123223333))
	fmt.Println(uintptr(a))
}

type data struct {
	value  uint32   // 4byte
	Label  [10]byte // 10byte
	Active bool     // 1byte
	// where is 1 byte???? 16바이트라매....
}

func byteFromData(d data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func bytesFromDataUnsafe(d data) [16]byte {
	if isLE {
		d.value = bits.ReverseBytes32(d.value)
	}
	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}

func dataFromBytes(b [16]byte) data {
	d := data{}
	d.value = binary.BigEndian.Uint32(b[4:])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

var isLE bool

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = xb[0] == 0x00
}

func dataFromBytesUnsafe(b [16]byte) data {
	data := *(*data)(unsafe.Pointer(&b))
	if isLE {
		data.value = bits.ReverseBytes32(data.value)
	}
	return data
}
