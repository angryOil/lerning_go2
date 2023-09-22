package main

var _ (ToString) = &name{}

type ToString interface {
	ToString() string
}

type name struct {
	firstName string
}

func (n name) ToString() string {
	//TODO implement me
	panic("implement me")
}
