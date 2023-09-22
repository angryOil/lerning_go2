package main

import "fmt"

func main() {
	useNilMap()
	useNotNilMap()
	useMapSlice()
	donKnowLength()
}

func useNilMap() {
	var nilMap map[int]int
	fmt.Println(nilMap == nil)
	fmt.Println(len(nilMap))
}

func useNotNilMap() {
	var notNilMap = map[int]int{}
	fmt.Println(notNilMap == nil)
	fmt.Println(len(notNilMap))
}

func useMapSlice() {
	teams := map[string][]string{
		"A": {"a", "b", "c"},
		"b": {"1", "2", "3"},
	}
	fmt.Println(teams)
}

func donKnowLength() {
	ages := make(map[int][]string)
	ages[0] = []string{"b", "c"}
	fmt.Println(ages)
	fmt.Println(len(ages))
	fmt.Println()
}
