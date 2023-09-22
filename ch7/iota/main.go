package main

import "errors"

func main() {

}

type mailate int

const (
	uncate mailate = iota
	personal
	spam
)

func validate(i mailate) error {
	if uncate == i {
		return errors.New("uncate is prediated")
	}

	return nil
}