package main

import "fmt"

func main() {
	node1 := newNode(1)
	node1.push(2).push(3).push(4)
	for node1 != nil {
		fmt.Print(node1.val, "-")
		node1 = node1.next
	}
}

type node[T any] struct {
	val  T
	next *node[T]
}

func newNode[T any](v T) *node[T] {
	return &node[T]{val: v}
}

func (n *node[T]) push(v T) *node[T] {
	node := newNode(v)
	n.next = node
	return node
}
