package double

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]

	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.tail == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.count++
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count() {
		return nil
	}
	i := 0
	for n := l.root; n != nil; n = n.next {
		if i == idx {
			return n
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) PushFront(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.root == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}

	l.root.prev = n
	l.root = n
	n.next = l.root
	l.count++
}

func (l *LinkedList[T]) InsertAfter(n *Node[T], val T) {
	if !l.IsIncluded(n) {
		return
	}

	newNode := &Node[T]{
		Value: val,
	}

	//if l.tail == n {
	//	tail := l.tail
	//	l.tail = n
	//	tail.next = n
	//	n.prev = tail
	//	l.count++
	//	return
	//}

	nextNode := n.next
	if nextNode != nil {
		nextNode.prev = newNode
	}
	if n == l.tail {
		l.tail = newNode
	}
	n.next = newNode
	newNode.prev = n
	newNode.next = nextNode
	l.count++
}

func (l *LinkedList[T]) IsIncluded(searchNode *Node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if n == searchNode {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(n *Node[T], val T) {
	if !l.IsIncluded(n) {
		return
	}

	newNode := &Node[T]{
		Value: val,
	}
	if n == l.root {
		l.root = newNode
		newNode.next = n
		n.prev = newNode
		l.count++
		return
	}

	prevNode := n.prev
	prevNode.next = newNode
	newNode.prev = prevNode
	newNode.next = n
	n.prev = newNode
	l.count++
	return
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	oldRoot := l.root
	if l.root == l.tail {
		l.root = nil
		l.tail = nil
		l.count--
		return oldRoot
	}
	newRoot := l.root.next
	l.root = newRoot
	newRoot.prev = nil
	l.count--
	return oldRoot
}
func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}
	oldTail := l.tail
	if l.root == l.tail {
		l.root = nil
		l.tail = nil
		l.count--
		return oldTail
	}

	newTail := l.tail.prev
	l.tail = newTail
	newTail.next = nil
	l.count--
	return oldTail
}

func (l *LinkedList[T]) Remove(targetNode *Node[T]) {
	if !l.IsIncluded(targetNode) {
		return
	}
	if l.root == l.tail {
		l.root = nil
		l.tail = nil
		l.count--
		return
	}
	if l.root == targetNode {
		l.PopFront()
		return
	} else if l.tail == targetNode {
		l.PopBack()
		return
	}
	prev := targetNode.prev
	next := targetNode.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	targetNode.prev = nil
	targetNode.next = nil
	l.count--
}

func (l *LinkedList[T]) Reverse() {
	if l.root == nil {
		return
	}
	target := l.root

	for target != nil {
		next := target.next
		prev := target.prev

		target.next = prev
		target.prev = next

		target = next
	}

	root := l.root
	tail := l.tail
	l.root = tail
	l.tail = root
}
