package main

import "fmt"

func main() {
	useSingleLinkedList()
	useDoubleLinkedList()
	useArr()
}

func useArr() {
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b [11]int
	copy(b[0:], a[0:5])
	b[5] = 100
	copy(b[6:], a[5:])
	fmt.Println(a)
	fmt.Println(b)
}
func useDoubleLinkedList() {
	root := &doubleLinkedNode[int]{nil, nil, 10}
	root.next = &doubleLinkedNode[int]{nil, root, 20}
	root.next.next = &doubleLinkedNode[int]{nil, root.next, 30}
	root.next.next.next = &doubleLinkedNode[int]{nil, root.next.next, 40}

	tail := root.next.next.next
	for n := tail; n != nil; n = n.prev {
		fmt.Println("double  val:", n.val)
	}
}

// 불연속 메모리
func useSingleLinkedList() {
	root := &singleNode[int]{nil, 10}
	root.next = &singleNode[int]{nil, 20}
	root.next.next = &singleNode[int]{nil, 30}

	for n := root; n != nil; n = n.next {
		fmt.Println("singleNode val:", n.val)
	}
}

type singleNode[T any] struct {
	next *singleNode[T]
	val  T
}

type doubleLinkedNode[T any] struct {
	next *doubleLinkedNode[T]
	prev *doubleLinkedNode[T]
	val  T
}
