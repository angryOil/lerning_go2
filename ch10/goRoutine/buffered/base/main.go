package main

import "fmt"

func main() {
	fmt.Println(processChannel())
}

func processChannel() []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func(val int) {
			results <- val * 22
		}(i)
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	close(results)
	return out
}
