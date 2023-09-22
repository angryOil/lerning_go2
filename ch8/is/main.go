package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func main() {
	useIs()
}

type myErr struct {
	codes []int
}

func (me myErr) Error() string {
	return fmt.Sprintf("codes:%v", me.codes)
}

func (me myErr) Is(target error) bool {
	var my myErr
	errors.As(me, &my)
	if me2, ok := target.(myErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

func useIs() {
	err := fileChecker("adkdsjfkladsjklads")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("os.ErrNotExist 에러입니다 ")
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in filechecker: %w", err)
	}
	f.Close()
	return nil
}
