package cmp

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_createPerson(t *testing.T) {
	expected := person{
		Name: "joy",
		Age:  22,
	}
	result := createPerson("joy", 22)
	//assert.Equal(t, result, expected)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

}
