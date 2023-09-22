package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := fileChecker("fgagfa")
	if err != nil {
		fmt.Println(err)
		if unwrapedErr := errors.Unwrap(err); unwrapedErr != nil {
			fmt.Println(unwrapedErr)
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

type status int
type statusErr struct {
	sta status
	msg string
	err error
}

func (se statusErr) Error() string {
	return se.msg

}

func (se statusErr) Unwrap() error {
	return se.err
}
