package ch13

import "time"

type person struct {
	Name      string
	Age       int
	DateAdded time.Time
}

func createPerson(name string, age int) person {
	return person{
		Name:      name,
		Age:       age,
		DateAdded: time.Now(),
	}
}
