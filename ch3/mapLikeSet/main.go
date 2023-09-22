package main

import "fmt"

func main() {
	useMapLikeSet()
}

func useMapLikeSet() {
	intSet := map[int]bool{}
	vals := []int{1, 4, 5, 3, 1, 1, 1, 9}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[50])
	fmt.Println(intSet[3])
}
