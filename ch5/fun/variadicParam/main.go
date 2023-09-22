package main

import (
	"fmt"
)

func main() {
	useAddTo()
}

func useAddTo() {
	nums := addTo(100, 1, 2, 3, 4, 5, 6)
	fmt.Println(nums)
	nums2 := addTo(10000, []int{1, 2, 3, 45, 1, 1, 1, 1, 234, 345, 56}...)
	fmt.Println(nums2)
	arr := []int{1, 2, 3, 3, 4}
	arr[2] = 10
	fmt.Println(arr)
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
