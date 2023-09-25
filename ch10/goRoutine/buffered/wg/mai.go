package main

import (
	"fmt"
	"sync"
)

// java stream Lists.of(1,2,3,).myMap(a -> b).myMap(b -> c)
func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	results := processAndGather(in, func(i int) int {
		fmt.Println(i)
		return i * i
	}, 3)
	fmt.Println(results)
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
			wg.Add(1)
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}
