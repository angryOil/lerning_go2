package ch13

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_createPerson(t *testing.T) {
	comparer := cmp.Comparer(func(x, y person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	expected := person{
		Name: "joy",
		Age:  22,
	}
	result := createPerson("joy", 22)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Errorf(diff)
	}

}
