package cmp

import "time"

type person struct {
	Name    string
	Age     int
	DateAdd time.Time
}

func createPerson(name string, age int) person {
	return person{
		Name:    name,
		Age:     age,
		DateAdd: time.Now(),
	}
}
