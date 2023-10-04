package ch10

import "fmt"

func buggedLoop() {
	a := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

func fixedLoopByGoParam() {
	a := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(i int) {
			ch <- i * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
func fixedLoopByShadowing() {
	a := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(a))
	for _, v := range a {
		v := v
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
