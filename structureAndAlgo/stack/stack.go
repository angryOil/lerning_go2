package stack

import "lerning_go2/structureAndAlgo/makeLinkedList/double"

type Stack[T any] struct {
	l *double.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		l: &double.LinkedList[T]{},
	}
}

func (s *Stack[T]) Push(val T) {
	s.l.PushBack(val)
}

func (s *Stack[T]) Pop() T {
	return s.l.PopBack().Value
}
