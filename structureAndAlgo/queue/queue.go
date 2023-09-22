package queue

import "lerning_go2/structureAndAlgo/makeLinkedList/double"

type Queue[T any] struct {
	l *double.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: &double.LinkedList[T]{},
	}
}

func (q *Queue[T]) Push(val T) {
	q.l.PushBack(val)
}

func (q *Queue[T]) Pop() T {
	return q.l.PopFront().Value
}
