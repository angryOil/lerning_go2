package double

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Nil(t, l.GetAt(4))
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.Count())

	l.PushFront(4)
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 4, l.Count())
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	n := l.GetAt(1)
	l.InsertAfter(n, 4)

	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(2).next.Value)
	assert.Equal(t, 4, l.GetAt(2).Value)

	tempNode := &Node[int]{
		Value: 1000,
	}
	l.InsertAfter(tempNode, 200)
	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(2).next.Value)
	assert.Equal(t, 4, l.GetAt(2).Value)

}

func Test_linkedList_insertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	n := l.GetAt(1)
	l.InsertBefore(n, 4)

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.GetAt(3).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 1000,
	}
	l.InsertAfter(tempNode, 200)
	assert.Equal(t, 4, l.Count())

	l.InsertAfter(l.GetAt(3), 10)
	assert.Equal(t, 5, l.Count())
	assert.Equal(t, 10, l.Back().Value)

	l.InsertBefore(l.Front(), 10)
	assert.Equal(t, 10, l.Front().Value)
}

func Test_linkedList_PopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.Equal(t, 3, l.Count())
	rootNext := l.root.next
	l.PopFront()
	assert.Equal(t, 2, l.root.Value)
	assert.Equal(t, 3, l.root.next.Value)
	assert.Nil(t, rootNext.prev)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 2, l.Count())

	l.PopFront()
	assert.Equal(t, 3, l.root.Value)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 1, l.count)

	l.PopFront()
	assert.Nil(t, l.root)
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
	assert.Equal(t, 0, l.count)

	l.PopFront()
	l.PopFront()
	l.PopFront()
	l.PopFront()
	assert.Nil(t, l.root)
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
	assert.Equal(t, 0, l.count)
}

func Test_linkedList_Remove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(l.GetAt(1))
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 3, l.GetAt(1).Value)

	l.Remove(l.GetAt(1))
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 1, l.Front().Value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

	l.Remove(l.GetAt(0))
	l.Remove(l.GetAt(0))
	l.Remove(l.GetAt(0))
	l.Remove(l.GetAt(0))
	l.Remove(l.GetAt(0))
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

}

func TestLinkedList_Reverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	l.Reverse()

	assert.Equal(t, 5, l.Front().Value)
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Equal(t, 2, l.GetAt(3).Value)
	assert.Equal(t, 1, l.Back().Value)
}
