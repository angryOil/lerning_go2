package main

import (
	"errors"
	"fmt"
)

func main() {
	useBoilerplateCode()
	useDefer()
}

func useDefer() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	fir, err := getError("2")
	sec, err := getError(fir)
	thr, err := getError(sec)
	fmt.Println(thr)
}
func useBoilerplateCode() {
	fir, err := getError("1")
	if err != nil {
		fmt.Println(err)
	}
	sec, err := getError(fir)
	if err != nil {
		fmt.Println(err)
	}
	thr, err := getError(sec)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(thr)
}

func getError(msg string) (string, error) {
	return msg, errors.New(msg)
}
