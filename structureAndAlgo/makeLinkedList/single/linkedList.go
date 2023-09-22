package single

type node[T any] struct {
	next  *node[T]
	value T
}

type linkedList[T any] struct {
	root *node[T]
	tail *node[T]

	count int
}

func (l *linkedList[T]) pushBack(value T) {
	node := &node[T]{
		value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	// 현재 마지막의 다음은 삽입된 노드
	l.tail.next = node
	// 테일도 마지막 노드로 변경
	l.tail = node
}

func (l *linkedList[T]) pushFront(value T) {
	node := &node[T]{
		value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root
	l.root = node
}

func (l *linkedList[T]) Front() *node[T] {
	return l.root
}

func (l *linkedList[T]) Back() *node[T] {
	return l.tail
}

func (l *linkedList[T]) Count() int {
	cnt := 0

	for n := l.root; n != nil; n = n.next {
		cnt++
	}
	return cnt
}

func (l *linkedList[T]) Count2() int {
	return l.count
}

func (l *linkedList[T]) GetAt(idx int) *node[T] {
	if idx >= l.Count2() {
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

func (l *linkedList[T]) InsertAfter(n *node[T], value T) {
	if !l.isIncluded(n) {
		return
	}
	newNode := &node[T]{
		value: value,
	}

	originNext := n.next
	n.next = newNode
	newNode.next = originNext

	// 위 아래는 같은 코드임
	// n.next, newNode.next = newNode, n.next
	l.count++
}

func (l *linkedList[T]) insertBefore(n *node[T], value T) {
	if n == l.root {
		l.pushFront(value)
		return
	}
	prevNode := l.findPrevNode(n)
	if prevNode == nil {
		return
	}

	newNode := &node[T]{
		value: value,
	}
	prevNode.next = newNode
	newNode.next = n
	l.count++
}

func (l *linkedList[T]) isIncluded(searchNode *node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if n == searchNode {
			return true
		}
	}
	return false
}

func (l *linkedList[T]) findPrevNode(searchNode *node[T]) *node[T] {
	for n := l.root; n != nil; n = n.next {
		if n.next == searchNode {
			return n
		}
	}
	return nil
}

func (l *linkedList[T]) PopFront() {
	if l.root == nil {
		return
	}

	//l.root.next, l.root = nil, l.root.next
	l.root = l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--
}

func (l *linkedList[T]) Remove(n *node[T]) {
	if n == l.root {
		l.PopFront()
		return
	}
	prev := l.findPrevNode(n)
	if prev == nil {
		return
	}

	prev.next = n.next
	n.next = nil
	if n == l.tail {
		l.tail = prev
	}

	l.count--
}

func (l *linkedList[T]) Reverse() {
	var reveredLinkedList linkedList[T]

	for l.root != nil {
		target := l.root
		reveredLinkedList.pushFront(target.value)
		l.PopFront()
	}

	*l = reveredLinkedList
}

func (l *linkedList[T]) Reverse2() {
	if l.root == nil {
		return
	}
	first := l.root
	target := l.root.next
	first.next = nil
	for target != nil {
		// 다음 타겟
		nextTarget := target.next

		target.next = first

		first = target

		target = nextTarget
	}

	l.root, l.tail = l.tail, l.root
}
