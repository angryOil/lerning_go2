package ch10

import "fmt"

func doDeadLock() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v1 := 1
		ch1 <- v1
		v2 := <-ch2
		fmt.Println("in go ", v1, v2)
	}()
	v1 := 2
	ch2 <- v1
	v2 := <-ch1
	fmt.Println("out go ", v1, v2)
}

func useSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v1 := 1
		ch1 <- v1
		fmt.Println("in go ", v1)
		v2 := <-ch2
		fmt.Println("in go ", v1, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
		fmt.Println("out go ch2 in ch2 <- v")
	case v2 = <-ch1:
		fmt.Println("out go in v2 = <-ch1 ", v, v2)
	}
	fmt.Println("out go out case", v, v2)
}
