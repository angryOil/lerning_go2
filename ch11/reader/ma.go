package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//useStrReader()
	useMyGZipReader()
	nowDir, _ := getNowDir()
	fmt.Println(nowDir)
}

func getNowDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func useMyGZipReader() {
	r, closer, err := buildGZipReader("ch11/reader/hello.gz")
	defer closer()
	if err != nil {
		fmt.Println("hello")
	}
	counts, err := countLetters(r)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(counts)
}

func useStrReader() {
	s := "the quick brown fox jumped over the lazy dog 한글입니다?1123"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(counts)
}

func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func countLetters(sr io.Reader) (map[string]int, error) {
	buf := make([]byte, 128)
	out := map[string]int{}
	for {
		n, err := sr.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
