package myMap

import (
	"golang.org/x/exp/constraints"
)

type sparseSet[TKey constraints.Ordered, TValue any] struct {
	dense  []element[TKey, TValue]
	sparse map[TKey]int
}

func newSparseSet[TKey constraints.Ordered, TValue any]() *sparseSet[TKey, TValue] {
	return &sparseSet[TKey, TValue]{
		sparse: make(map[TKey]int),
	}
}

func (s *sparseSet[TKey, TValue]) Add(key TKey, value TValue) {
	if idx, ok := s.sparse[key]; ok {
		s.dense[idx].value = value
		return
	}

	s.dense = append(s.dense, element[TKey, TValue]{
		key:   key,
		value: value,
	})
	s.sparse[key] = len(s.dense) - 1
}

func (s *sparseSet[TKey, TValue]) Get(key TKey) (TValue, bool) {
	if idx, ok := s.sparse[key]; ok {
		return s.dense[idx].value, true
	}
	var notFound TValue
	return notFound, false
}
