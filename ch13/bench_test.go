package ch13

import (
	"fmt"
	"testing"
)

var testData = "testdata/data.txt"
var dataSize = 365568

func Test_fileLen(t *testing.T) {
	result, err := fileLen(testData, 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != dataSize {
		t.Error("expected ", dataSize, " but", result)
	}
}

var blackHole int

func Benchmark_fileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := fileLen(testData, 1)
		if err != nil {
			b.Fatal(err)
		}
		blackHole = result
	}
}

func Benchmark_fileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("fileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := fileLen(testData, v)
				if err != nil {
					b.Fatal(err)
				}
				blackHole = result
			}
		})
	}
}
func Benchmark_fileLen2(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("fileLen-%d", v), func(b *testing.B) {
			buf := make([]byte, v)
			for i := 0; i < b.N; i++ {
				result, err := fileLen2(testData, buf)
				if err != nil {
					b.Fatal(err)
				}
				blackHole = result
			}
		})
	}
}
