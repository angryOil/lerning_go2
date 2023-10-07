package ch13

import "os"

func fileLen(f string, bufSize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	buf := make([]byte, bufSize)
	for {
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
}

func fileLen2(f string, buf []byte) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	for {
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
}
