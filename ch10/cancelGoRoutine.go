package ch10

import "fmt"

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
				fmt.Println("done")
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}
