package myMap

import "hash/crc32"

// hash 구조는 spase 한 구조라서 캐쉬 미스가 잘일어남
// ex) ["world","??","","abc"] << 이와같이 key 크기와 관계없이 해쉬가 일어나고 뜸성뜸성 있기때문에
// 캐쉬에 친화적이지 않음(대신 탐색,추가,삭제 빅오는 O(1) )
// const arrSize = 3571
const arrSize = 3

type hashData[T any] struct {
	key   string
	value T
}

type hashMap[T any] struct {
	arr [arrSize][]hashData[T]
}

func hashFn(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (h *hashMap[T]) Add(key string, value T) {
	hash := hashFn(key)
	var hd = hashData[T]{
		key:   key,
		value: value,
	}
	h.arr[hash%arrSize] = append(h.arr[hash%arrSize], hd)
}

func (h *hashMap[T]) Get(key string) (T, bool) {
	hash := hashFn(key)

	for _, hd := range h.arr[hash%arrSize] {
		if hd.key == key {
			return hd.value, true
		}
	}
	//val := h.arr[hash%arrSize]
	//return val.value, key == val.key
	var tempT T
	return tempT, false
}
