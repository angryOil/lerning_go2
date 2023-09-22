package main

import "fmt"

func main() {
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}
	defer cancel()
	a, ok := <-ch
	fmt.Println(a, ok)
	a, ok = <-ch
	fmt.Println(a, ok)
}

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				fmt.Println("i?", i)
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}
