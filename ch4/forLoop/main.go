package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8}
	fmt.Println(evenVals)
	for i, v := range evenVals {
		evenVals[i] = v * 2
	}
	fmt.Println(evenVals)
}
