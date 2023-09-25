package myMap

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// 이진트리 사용
// 삽입 O(N) 위치 찾기는 logN 이지만 갯수만큼 뒤로 밀어야함
// 조회 O(logN) 이진트리 탐색
// 삭제 O(N) 탐색 LogN + 삭제후 정렬 n 개만큼
// 정렬 상태 유지하고 싶을경우 사용
// 캐쉬 친화적 dense(뺴곡함)
type element[TKey constraints.Ordered, TValue any] struct {
	key   TKey
	value TValue
}

type sortedMap[TKey constraints.Ordered, TValue any] struct {
	Arr []element[TKey, TValue]
}

func (s *sortedMap[TKey, TValue]) Add(key TKey, value TValue) {

	// sort.Search bool 조건이 되는 최소값을 찾아줌
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].key == key {
		s.Arr[idx].value = value
		return
	}

	s.Arr = append(s.Arr[:idx],
		append([]element[TKey, TValue]{
			{key: key, value: value},
		}, s.Arr[idx:]...)...)
}

func (s *sortedMap[TKey, TValue]) Get(key TKey) (TValue, bool) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].key == key {
		return s.Arr[idx].value, true
	}
	var empty TValue
	return empty, false
}

func (s *sortedMap[TKey, TValue]) Remove(key TKey) bool {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].key == key {
		s.Arr = append(s.Arr[:idx], s.Arr[idx+1:]...)
		return true
	}
	return false
}
