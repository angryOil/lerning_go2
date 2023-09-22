package single

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushBack(t *testing.T) {
	var l linkedList[int]
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.pushBack(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	l.pushBack(2)
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 2, l.Back().value)

	l.pushBack(3)
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 3, l.Back().value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	assert.Equal(t, 1, l.GetAt(0).value)
	assert.Equal(t, 2, l.GetAt(1).value)
	assert.Equal(t, 3, l.GetAt(2).value)
	assert.Nil(t, l.GetAt(4))
}

func TestPushFront(t *testing.T) {
	var l linkedList[int]
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.pushFront(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	l.pushFront(2)
	assert.Equal(t, 2, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	l.pushFront(3)
	assert.Equal(t, 3, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())
}

func TestInsertAfter(t *testing.T) {
	var l linkedList[int]

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)

	n := l.GetAt(1)
	l.InsertAfter(n, 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).value)
	assert.Equal(t, 2, l.GetAt(1).value)
	assert.Equal(t, 3, l.Back().value)

	tempNode := &node[int]{
		value: 1000,
	}
	l.InsertAfter(tempNode, 200)
	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())

}

func Test_linkedList_insertBefore(t *testing.T) {
	var l linkedList[int]

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)

	n := l.GetAt(1)
	l.insertBefore(n, 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 1, l.GetAt(0).value)
	assert.Equal(t, 4, l.GetAt(1).value)
	assert.Equal(t, 2, l.GetAt(2).value)
	assert.Equal(t, 3, l.GetAt(3).value)
	assert.Equal(t, 3, l.Back().value)

	tempNode := &node[int]{
		value: 1000,
	}
	l.InsertAfter(tempNode, 200)
	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())
}
func Test_linkedList_insertBefore_root(t *testing.T) {
	var l linkedList[int]

	l.pushBack(1)
	n := l.GetAt(0)
	l.insertBefore(n, 4)

	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 4, l.GetAt(0).value)
	assert.Equal(t, 1, l.GetAt(1).value)
	assert.Equal(t, 1, l.Back().value)

	tempNode := &node[int]{
		value: 1000,
	}
	l.InsertAfter(tempNode, 200)
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())
}

func Test_linkedList_PopFront(t *testing.T) {
	var l linkedList[int]

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)

	l.PopFront()

	assert.Equal(t, 2, l.root.value)
	assert.Equal(t, 3, l.root.next.value)
	assert.Equal(t, 2, l.Front().value)
	assert.Equal(t, 3, l.Back().value)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())

	l.PopFront()

	assert.Equal(t, 3, l.root.value)
	assert.Equal(t, 3, l.Front().value)
	assert.Equal(t, 3, l.Back().value)
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
	var l linkedList[int]

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)

	l.Remove(l.GetAt(1))
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 1, l.GetAt(0).value)
	assert.Equal(t, 3, l.GetAt(1).value)

	l.Remove(l.GetAt(1))
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 1, l.Back().value)
	assert.Equal(t, 1, l.Front().value)

}

func TestLinkedList_Reverse(t *testing.T) {
	var l linkedList[int]

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)
	l.pushBack(4)
	l.pushBack(5)

	l.Reverse()

	assert.Equal(t, 5, l.Count2())
	assert.Equal(t, 5, l.Front().value)
	assert.Equal(t, 4, l.GetAt(1).value)
	assert.Equal(t, 3, l.GetAt(2).value)
	assert.Equal(t, 2, l.GetAt(3).value)
	assert.Equal(t, 1, l.Back().value)
}
func TestLinkedList_Reverse2(t *testing.T) {
	var l linkedList[int]

	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)
	l.pushBack(4)
	l.pushBack(5)

	l.Reverse2()

	assert.Equal(t, 5, l.Count2())
	assert.Equal(t, 5, l.Front().value)
	assert.Equal(t, 4, l.GetAt(1).value)
	assert.Equal(t, 3, l.GetAt(2).value)
	assert.Equal(t, 2, l.GetAt(3).value)
	assert.Equal(t, 1, l.Back().value)
}
